FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o contact-app .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache \
    ca-certificates \
    postgresql17-client

COPY --from=builder /app/contact-app .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]