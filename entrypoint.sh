#!/bin/sh

set -e

echo "🚀 Starting Contact List CLI..."

docker-entrypoint.sh postgres &

export DATABASE_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"

echo "⏳ Waiting for PostgreSQL..."

until pg_isready \
    -h localhost \
    -p 5432 \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB"
do
    sleep 2
done

echo "✅ PostgreSQL is ready"

exec ./contact-app