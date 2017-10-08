#!/bin/bash

echo "> Running server tests..."

go test -v ./server/pkg/...
