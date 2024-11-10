FROM golang:1.23.1-alpine AS api

WORKDIR /builder
COPY ./server .
RUN go build -o site-manager-api .

FROM node:22-alpine AS frontend

WORKDIR /builder
COPY ./client .
RUN corepack enable
RUN yarn install
RUN yarn build

FROM nginx:alpine AS final
WORKDIR /app
COPY --from=api /builder/site-manager-api /app
COPY --from=frontend /builder/dist /app
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

ENTRYPOINT ["/bin/sh", "-c", "nginx && ./site-manager-api -run"]