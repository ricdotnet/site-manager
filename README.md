#### Development

- Install golang
- Get a mysql database
- Copy `.env.example` to `.env.development`

#### Run

```sh
go run . -run --env-file .env.development
```

Any database changes should be applied automatically. If it is the first time running, tables should be created.