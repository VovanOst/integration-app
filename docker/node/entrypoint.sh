#!/bin/sh
set -e

echo "Installing dependencies..."
npm install

echo "Running development server..."
npm run dev -- --host 0.0.0.0