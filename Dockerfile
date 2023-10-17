FROM golang:1.20 as builder
WORKDIR /builder
COPY . .
RUN go get
RUN go build -o server .

FROM builder as prod
WORKDIR /production
COPY --from=builder /builder/.env /production
COPY --from=builder /builder/server /production

EXPOSE 3000
CMD ["./server", "-run"]