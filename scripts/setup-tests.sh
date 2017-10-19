#!/bin/bash

echo "> Removing database file if it exists..."

if [ -f "database.db" ]; then
  rm "database.db"
fi
