FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8080

CMD ["./app"]
