#!/bin/bash

mkdir -p certs
cd certs

if [[ -z "${GOROOT}" ]]; then
	echo "GOROOT is not defined. Set GOROOT to continue..."
	exit 1
fi

echo "Generating certs..."

go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
