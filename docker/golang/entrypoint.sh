#!/bin/sh
set -e

echo "Waiting for database..."
until nc -z db 5432; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Running Go application..."
./api
