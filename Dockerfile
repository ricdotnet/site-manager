FROM golang:1.20 AS api
WORKDIR /builder
COPY . .
RUN go get
RUN go build -o server .

FROM node:18 AS client
COPY --from=api /builder/server /production/server
COPY --from=api /builder/.env /production/.env
COPY --from=api /builder/boot.sh /production/boot.sh
WORKDIR /builder
COPY . .
WORKDIR /builder/client
RUN yarn install
RUN yarn build

FROM nginx
COPY --from=client /production/server /production/server
COPY --from=client /production/.env /production/.env
COPY --from=client /production/boot.sh /production/boot.sh
COPY --from=client /builder/client/dist /production/client
COPY ./client/nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3001
EXPOSE 80
WORKDIR /production
CMD ["/bin/bash", "boot.sh"]