#!/bin/sh

set -e

echo "🚀 Starting Contact List CLI..."

env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /var/app/.env

echo "⏳ Waiting for PostgreSQL..."


until pg_isready -d "$DATABASE_URL" > /dev/null 2>&1
do
    sleep 2
done

echo "✅ PostgreSQL is ready"

exec ./contact-app