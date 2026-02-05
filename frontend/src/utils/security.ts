/**
 * Security utility class - prevent XSS attacks
 */

import DOMPurify from 'dompurify';

// Configure DOMPurify security policy
const DOMPurifyConfig = {
  // Allowed tags
  ALLOWED_TAGS: [
    'p', 'br', 'strong', 'em', 'u', 's', 'del', 'ins',
    'h1', 'h2', 'h3', 'h4', 'h5', 'h6',
    'ul', 'ol', 'li', 'blockquote', 'pre', 'code',
    'a', 'img', 'table', 'thead', 'tbody', 'tr', 'th', 'td',
    'div', 'span', 'figure', 'figcaption', 'think'
  ],
  // Allowed attributes
  ALLOWED_ATTR: [
    'href', 'title', 'alt', 'src', 'class', 'id', 'style',
    'target', 'rel', 'width', 'height'
  ],
  // Allowed protocols
  ALLOWED_URI_REGEXP: /^(?:(?:(?:f|ht)tps?|mailto|tel|callto|cid|xmpp):|[^a-z]|[a-z+.\-]+(?:[^a-z+.\-:]|$))/i,
  // Forbidden tags and attributes
  FORBID_TAGS: ['script', 'object', 'embed', 'form', 'input', 'button'],
  FORBID_ATTR: ['onerror', 'onload', 'onclick', 'onmouseover', 'onfocus', 'onblur'],
  // Other security configurations
  KEEP_CONTENT: true,
  RETURN_DOM: false,
  RETURN_DOM_FRAGMENT: false,
  RETURN_DOM_IMPORT: false,
  SANITIZE_DOM: true,
  SANITIZE_NAMED_PROPS: true,
  WHOLE_DOCUMENT: false,
  // Custom hook functions
  HOOKS: {
    // Process before sanitization
    beforeSanitizeElements: (currentNode: Element) => {
      // Remove all script tags
      if (currentNode.tagName === 'SCRIPT') {
        currentNode.remove();
        return null;
      }
      // Remove all event handlers
      const eventAttrs = ['onclick', 'onload', 'onerror', 'onmouseover', 'onfocus', 'onblur'];
      eventAttrs.forEach(attr => {
        if (currentNode.hasAttribute(attr)) {
          currentNode.removeAttribute(attr);
        }
      });
    },
    // Process after sanitization
    afterSanitizeElements: (currentNode: Element) => {
      // Ensure all links have rel="noopener noreferrer"
      if (currentNode.tagName === 'A') {
        const href = currentNode.getAttribute('href');
        if (href && href.startsWith('http')) {
          currentNode.setAttribute('rel', 'noopener noreferrer');
          currentNode.setAttribute('target', '_blank');
        }
      }
      // Ensure all images have alt attribute
      if (currentNode.tagName === 'IMG') {
        if (!currentNode.getAttribute('alt')) {
          currentNode.setAttribute('alt', '');
        }
      }
    }
  }
};

/**
 * Safely sanitize HTML content
 * @param html HTML string to sanitize
 * @returns Sanitized safe HTML string
 */
export function sanitizeHTML(html: string): string {
  if (!html || typeof html !== 'string') {
    return '';
  }
  
  try {
    return DOMPurify.sanitize(html, DOMPurifyConfig);
  } catch (error) {
    console.error('HTML sanitization failed:', error);
    // If sanitization fails, return escaped plain text
    return escapeHTML(html);
  }
}

/**
 * Escape HTML special characters
 * @param text Text to escape
 * @returns Escaped text
 */
export function escapeHTML(text: string): string {
  if (!text || typeof text !== 'string') {
    return '';
  }
  
  const map: { [key: string]: string } = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#x27;',
    '/': '&#x2F;',
    '`': '&#x60;',
    '=': '&#x3D;'
  };
  
  return text.replace(/[&<>"'`=\/]/g, (s) => map[s]);
}

/**
 * Validate if URL is safe
 * @param url URL to validate
 * @returns Whether URL is safe
 */
export function isValidURL(url: string): boolean {
  if (!url || typeof url !== 'string') {
    return false;
  }
  
  try {
    const urlObj = new URL(url);
    // Only allow http and https protocols
    return ['http:', 'https:'].includes(urlObj.protocol);
  } catch {
    return false;
  }
}

/**
 * Safely process Markdown content
 * @param markdown Markdown text
 * @returns Safe HTML string
 */
export function safeMarkdownToHTML(markdown: string): string {
  if (!markdown || typeof markdown !== 'string') {
    return '';
  }
  
  // First escape possible HTML tags
  const escapedMarkdown = markdown
    .replace(/<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi, '')
    .replace(/<iframe\b[^<]*(?:(?!<\/iframe>)<[^<]*)*<\/iframe>/gi, '')
    .replace(/<object\b[^<]*(?:(?!<\/object>)<[^<]*)*<\/object>/gi, '')
    .replace(/<embed\b[^<]*(?:(?!<\/embed>)<[^<]*)*<\/embed>/gi, '');
  
  return escapedMarkdown;
}

/**
 * Clean user input
 * @param input User input
 * @returns Cleaned safe input
 */
export function sanitizeUserInput(input: string): string {
  if (!input || typeof input !== 'string') {
    return '';
  }
  
  // Remove control characters
  let cleaned = input.replace(/[\x00-\x1F\x7F-\x9F]/g, '');
  
  // Limit length
  if (cleaned.length > 10000) {
    cleaned = cleaned.substring(0, 10000);
  }
  
  return cleaned.trim();
}

/**
 * Validate if image URL is safe
 * @param url Image URL
 * @returns Whether image URL is safe
 */
export function isValidImageURL(url: string): boolean {
  if (!isValidURL(url)) {
    return false;
  }
  
  return true;
}

/**
 * Create safe image element
 * @param src Image source
 * @param alt Alternative text
 * @param title Title
 * @returns Safe image HTML
 */
export function createSafeImage(src: string, alt: string = '', title: string = ''): string {
  if (!isValidImageURL(src)) {
    return '';
  }
  
  const safeSrc = escapeHTML(src);
  const safeAlt = escapeHTML(alt);
  const safeTitle = escapeHTML(title);
  
  return `<img src="${safeSrc}" alt="${safeAlt}" title="${safeTitle}" class="markdown-image" style="max-width: 100%; height: auto;">`;
}
