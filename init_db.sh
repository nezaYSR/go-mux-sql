#!/bin/bash

# Check if the database exists
if echo "SHOW DATABASES;" | mysql -u root -proot | grep -q "scroll_table"; then
  echo "Database already exists"
else
  # Create the database
  echo "Creating database..."
  echo "CREATE DATABASE scroll_table;" | mysql -u root -proot
  echo "Database created"
fi

echo "Running main.go..."
cd cmd
go run main.go

