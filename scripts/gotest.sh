#!/usr/bin/env bash

echo "==> Running go test and creating a coverage profile..."

go test ./... -v -race -cover -coverprofile "coverage.all.out"

if [ $? -eq 1 ]; then
	exit 1
fi

cat coverage.all.out | grep -v '\.gen\.go' | grep -v 'mock_' > coverage.out
rm -f coverage.all.out

exit 0
