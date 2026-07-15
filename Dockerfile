FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o contact-app .

FROM postgres:alpine

ENV POSTGRES_PASSWORD=1

COPY lib/init.sql /docker-entrypoint-initdb.d/

WORKDIR /app

COPY --from=builder /app/contact-app .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]