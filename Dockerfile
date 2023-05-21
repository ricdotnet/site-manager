FROM golang:1.20

WORKDIR /app

COPY . .

RUN go get

RUN go build -o bin .

EXPOSE 3000

CMD ["./bin", "-run"]