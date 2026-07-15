#!/bin/sh

set -e

echo "🚀 Starting Contact List CLI..."

echo "⏳ Waiting for PostgreSQL..."

docker-entrypoint.sh postgres &
env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > .env

until pg_isready -q
do
    sleep 2
done

echo "✅ PostgreSQL is ready"

exec ./contact-app