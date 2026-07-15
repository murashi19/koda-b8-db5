#!/bin/sh

set -e

echo "🚀 Starting Contact List CLI..."

if [ -z "$DATABASE_URL" ]; then
    echo "❌ DATABASE_URL is not set"
    exit 1
fi

echo "⏳ Waiting for PostgreSQL..."

until pg_isready -d "$DATABASE_URL" > /dev/null 2>&1
do
    sleep 2
done

echo "✅ PostgreSQL is ready"

exec ./contact-app