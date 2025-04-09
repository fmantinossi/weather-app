FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./cmd/server/main.go

FROM debian:bullseye-slim

WORKDIR /

COPY --from=builder /app/app .

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

EXPOSE 8080

CMD ["./app"]