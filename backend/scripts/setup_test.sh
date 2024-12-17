#!/bin/bash

# Exit on error
set -e

echo "Setting up test environment..."

# Create test database
echo "Creating test database..."
mysql -u root -p < scripts/setup_test_db.sql

# Create test uploads directory
echo "Creating test uploads directory..."
mkdir -p ./test_uploads

echo "Test environment setup complete!"
