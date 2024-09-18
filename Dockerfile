FROM golang:1.23.0 AS builder
WORKDIR /builder
COPY . .
RUN go get \
&& go build -o server .

FROM builder AS prod
WORKDIR /production
COPY --from=builder /builder/.env /production
COPY --from=builder /builder/server /production

EXPOSE 3000
CMD ["./server", "-run"]