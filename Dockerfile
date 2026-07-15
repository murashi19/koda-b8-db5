FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o contact-app .

FROM postgres:alpine

ENV POSTGRES_USER=postgres2
ENV POSTGRES_PASSWORD=1
ENV POSTGRES_DB=postgres3


WORKDIR /app

COPY lib/init.sql /docker-entrypoint-initdb.d/
COPY --from=builder /app/contact-app .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]