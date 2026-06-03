FROM node:18-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.27-alpine
COPY --from=build /app/build /usr/share/nginx/html
RUN <<'EOF'
cat > /docker-entrypoint.d/40-env.sh <<'SCRIPT'
#!/bin/sh
set -e
: "${API_URL:=http://localhost:8080}"
echo "window.__API_URL__ = \"${API_URL}\";" > /usr/share/nginx/html/env.js
SCRIPT
chmod +x /docker-entrypoint.d/40-env.sh
EOF

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
