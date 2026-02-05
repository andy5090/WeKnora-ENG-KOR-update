#!/bin/sh

# Generate runtime configuration file, inject environment variables to frontend
cat > /usr/share/nginx/html/config.js << EOF
window.__RUNTIME_CONFIG__ = {
  MAX_FILE_SIZE_MB: ${MAX_FILE_SIZE_MB:-50}
};
EOF

# Process nginx configuration
export MAX_FILE_SIZE=${MAX_FILE_SIZE_MB}M
envsubst '${MAX_FILE_SIZE}' < /etc/nginx/templates/default.conf.template > /etc/nginx/conf.d/default.conf

# Start nginx
exec nginx -g 'daemon off;'
