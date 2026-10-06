#!/usr/bin/env bash
set -euo pipefail

echo "==> Checking mocks are up to date..."

go run github.com/vektra/mockery/v2@"${VERSION}"

if [ -n "$(git diff --stat)" ]; then
    echo "Mocks are outdated!"
    exit 1
fi

exit 0
