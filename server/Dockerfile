FROM golang:1.23.1-alpine3.20 AS builder

WORKDIR /builder

COPY ./server .

RUN go build -o dist .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /builder/dist /app

EXPOSE 3000

ENTRYPOINT ["./dist", "-run"]