// src/utils/request.js
import axios from "axios";
import { generateRandomString } from "./index";

// API base URL
const BASE_URL = import.meta.env.VITE_IS_DOCKER ? "" : "http://localhost:8080";


// Create Axios instance
const instance = axios.create({
  baseURL: BASE_URL, // Use configured API base URL
  timeout: 30000, // Request timeout
  headers: {
    "Content-Type": "application/json",
    "X-Request-ID": `${generateRandomString(12)}`,
  },
});


instance.interceptors.request.use(
  (config) => {
    // Add JWT token authentication
    const token = localStorage.getItem('weknora_token');
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    
    // Add cross-tenant access request header (if another tenant is selected)
    const selectedTenantId = localStorage.getItem('weknora_selected_tenant_id');
    const defaultTenantId = localStorage.getItem('weknora_tenant');
    if (selectedTenantId) {
      try {
        const defaultTenant = defaultTenantId ? JSON.parse(defaultTenantId) : null;
        const defaultId = defaultTenant?.id ? String(defaultTenant.id) : null;
        // If selected tenant ID differs from default tenant ID, add request header
        if (selectedTenantId !== defaultId) {
          config.headers["X-Tenant-ID"] = selectedTenantId;
        }
      } catch (e) {
        console.error('Failed to parse tenant info', e);
      }
    }
    
    config.headers["X-Request-ID"] = `${generateRandomString(12)}`;
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Token refresh flag to prevent multiple requests from refreshing token simultaneously
let isRefreshing = false;
let failedQueue: Array<{ resolve: Function; reject: Function }> = [];
let hasRedirectedOn401 = false;

// Process queued requests
const processQueue = (error: any, token: string | null = null) => {
  failedQueue.forEach(({ resolve, reject }) => {
    if (error) {
      reject(error);
    } else {
      resolve(token);
    }
  });
  
  failedQueue = [];
};

instance.interceptors.response.use(
  (response) => {
    // Handle logic based on business status code
    const { status, data } = response;
    if (status === 200 || status === 201) {
      return data;
    } else {
      return Promise.reject(data);
    }
  },
  async (error: any) => {
    const originalRequest = error.config;
    
    if (!error.response) {
      return Promise.reject({ message: "Network error, please check your network connection" });
    }
    
    // If it's a 401 from login endpoint, return error directly for toast display, don't redirect
    if (error.response.status === 401 && originalRequest?.url?.includes('/auth/login')) {
      const { status, data } = error.response;
      return Promise.reject({ status, message: (typeof data === 'object' ? data?.message : data) || 'Incorrect username or password' });
    }

    // If it's a 401 error and not a refresh token request, try to refresh token
    if (error.response.status === 401 && !originalRequest._retry && !originalRequest.url?.includes('/auth/refresh')) {
      if (isRefreshing) {
        // If token is being refreshed, add request to queue
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        }).then(token => {
          originalRequest.headers['Authorization'] = 'Bearer ' + token;
          return instance(originalRequest);
        }).catch(err => {
          return Promise.reject(err);
        });
      }
      
      originalRequest._retry = true;
      isRefreshing = true;
      
      const refreshToken = localStorage.getItem('weknora_refresh_token');
      
      if (refreshToken) {
        try {
          // Dynamically import refresh token API
          const { refreshToken: refreshTokenAPI } = await import('../api/auth/index');
          const response = await refreshTokenAPI(refreshToken);
          
          if (response.success && response.data) {
            const { token, refreshToken: newRefreshToken } = response.data;
            
            // Update token in localStorage
            localStorage.setItem('weknora_token', token);
            localStorage.setItem('weknora_refresh_token', newRefreshToken);
            
            // Update request header
            originalRequest.headers['Authorization'] = 'Bearer ' + token;
            
            // Process queued requests
            processQueue(null, token);
            
            return instance(originalRequest);
          } else {
            throw new Error(response.message || 'Token refresh failed');
          }
        } catch (refreshError) {
          // Refresh failed, clear all tokens and redirect to login page
          localStorage.removeItem('weknora_token');
          localStorage.removeItem('weknora_refresh_token');
          localStorage.removeItem('weknora_user');
          localStorage.removeItem('weknora_tenant');
          
          processQueue(refreshError, null);
          
          // Redirect to login page
          if (!hasRedirectedOn401 && typeof window !== 'undefined') {
            hasRedirectedOn401 = true;
            window.location.href = '/login';
          }
          
          return Promise.reject(refreshError);
        } finally {
          isRefreshing = false;
        }
      } else {
        // No refresh token, directly redirect to login page
        localStorage.removeItem('weknora_token');
        localStorage.removeItem('weknora_user');
        localStorage.removeItem('weknora_tenant');
        
        if (!hasRedirectedOn401 && typeof window !== 'undefined') {
          hasRedirectedOn401 = true;
          window.location.href = '/login';
        }
        
        return Promise.reject({ message: 'Please login again' });
      }
    }
    
    // Handle Nginx 413 Request Entity Too Large
    if (error.response.status === 413) {
      return Promise.reject({ 
        status: 413, 
        message: 'File size exceeds limit, please upload a smaller file',
        success: false
      });
    }

    const { status, data } = error.response;
    // Throw HTTP status code together for upper layer to judge 401 and other scenarios
    // Backend return format: { success: false, error: { code, message, details } }
    // Extract error.message as top-level message for frontend to use error?.message
    const errorMessage = typeof data === 'object' && data?.error?.message 
      ? data.error.message 
      : (typeof data === 'object' ? data?.message : data);
    return Promise.reject({ 
      status, 
      message: errorMessage,
      ...(typeof data === 'object' ? data : {}) 
    });
  }
);

export function get(url: string) {
  return instance.get(url);
}

export async function getDown(url: string) {
  let res = await instance.get(url, {
    responseType: "blob",
  });
  return res
}

export function postUpload(url: string, data = {}, onUploadProgress?: (progressEvent: any) => void) {
  return instance.post(url, data, {
    headers: {
      "Content-Type": "multipart/form-data",
      "X-Request-ID": `${generateRandomString(12)}`,
    },
    onUploadProgress,
  });
}

export function postChat(url: string, data = {}) {
  return instance.post(url, data, {
    headers: {
      "Content-Type": "text/event-stream;charset=utf-8",
      "X-Request-ID": `${generateRandomString(12)}`,
    },
  });
}

export function post(url: string, data = {}, config?: any) {
  return instance.post(url, data, config);
}

export function put(url: string, data = {}) {
  return instance.put(url, data);
}

export function del(url: string, data?: any) {
  return instance.delete(url, { data });
}
