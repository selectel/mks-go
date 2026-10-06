#!/usr/bin/env bash
set -uo pipefail

echo "==> Running golangci-lint..."

exit_code=0
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@"${VERSION}" run ./... || exit_code=$?

if [ "${exit_code}" -ne 0 ]; then
    echo ""
    echo "Golangci-lint found suspicious constructs. Please check the reported"; \
    echo "constructs and fix them if necessary before submitting the code for review."; \
    exit 1
fi

exit 0
