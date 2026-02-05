import { MessagePlugin } from "tdesign-vue-next";

// Declare global runtime configuration type
declare global {
  interface Window {
    __RUNTIME_CONFIG__?: {
      MAX_FILE_SIZE_MB?: number;
    };
  }
}

// Get maximum file size (MB) from runtime configuration, supports dynamic configuration in Docker environment
// Priority: runtime configuration > build-time environment variable > default value 50MB
const MAX_FILE_SIZE_MB = window.__RUNTIME_CONFIG__?.MAX_FILE_SIZE_MB 
  || Number(import.meta.env.VITE_MAX_FILE_SIZE_MB) 
  || 50;
const MAX_FILE_SIZE_BYTES = MAX_FILE_SIZE_MB * 1024 * 1024;

export function generateRandomString(length: number) {
  let result = "";
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

export function formatStringDate(date: any) {
  let data = new Date(date);
  let year = data.getFullYear();
  let month = String(data.getMonth() + 1).padStart(2, '0');
  let day = String(data.getDate()).padStart(2, '0');
  let hour = String(data.getHours()).padStart(2, '0');
  let minute = String(data.getMinutes()).padStart(2, '0');
  let second = String(data.getSeconds()).padStart(2, '0');
  return (
    year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second
  );
}
export function kbFileTypeVerification(file: any, silent = false) {
  let validTypes = ["pdf", "txt", "md", "docx", "doc", "jpg", "jpeg", "png", "csv", "xlsx", "xls"];
  let type = file.name.substring(file.name.lastIndexOf(".") + 1);
  if (!validTypes.includes(type)) {
    if (!silent) {
      MessagePlugin.error("Invalid file type!");
    }
    return true;
  }
  if (
    (type == "pdf" || type == "docx" || type == "doc") &&
    file.size > MAX_FILE_SIZE_BYTES
  ) {
    if (!silent) {
      MessagePlugin.error(`PDF/DOC files cannot exceed ${MAX_FILE_SIZE_MB}MB!`);
    }
    return true;
  }
  if ((type == "txt" || type == "md") && file.size > MAX_FILE_SIZE_BYTES) {
    if (!silent) {
      MessagePlugin.error(`TXT/MD files cannot exceed ${MAX_FILE_SIZE_MB}MB!`);
    }
    return true;
  }
  return false
}
