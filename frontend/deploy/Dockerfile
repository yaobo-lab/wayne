FROM node:11.1-alpine as build-stage

WORKDIR /app

COPY package*.json  ./
RUN npm config set registry https://registry.npm.taobao.org
RUN npm install
COPY . .
RUN npm run build:aot
RUN ls
RUN ls /app/dist/

# build server
FROM openresty/openresty:1.15.8.1-1-centos  as production-stage

COPY --from=build-stage /app/dist/ /usr/local/openresty/nginx/html/

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    sed -i '/index  index.html index.htm;/a\        try_files $uri $uri/ /index.html;' /etc/nginx/conf.d/default.conf

CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]
