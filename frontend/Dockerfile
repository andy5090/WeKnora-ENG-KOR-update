# Build stage
FROM node:24-alpine as build-stage

WORKDIR /app

# Set environment variables, ignore type check errors
ENV NODE_OPTIONS="--max-old-space-size=4096"
ENV VITE_IS_DOCKER=true

# Copy dependency files
COPY package*.json ./
COPY packages/xlsx-0.20.2.tgz ./packages/xlsx-0.20.2.tgz

# Install dependencies
RUN corepack enable
RUN pnpm install

# Copy project files
COPY . .

# Build application
RUN pnpm run build

# Production stage
FROM nginx:stable-alpine as production-stage

# Copy build artifacts to nginx service directory
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Copy nginx configuration template file
COPY nginx.conf /etc/nginx/templates/default.conf.template

# Copy startup script
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

# Set default environment variable (MB)
ENV MAX_FILE_SIZE_MB=50

# Expose port
EXPOSE 80

ENTRYPOINT ["/docker-entrypoint.sh"] 