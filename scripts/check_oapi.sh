#!/usr/bin/env bash
set -euo pipefail

echo "==> Checking oapi generated code is up to date..."

cd swagger
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@"${VERSION}" --config=oapi-codegen-models.yaml managed-kubernetes.swagger.yaml
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@"${VERSION}" --config=oapi-codegen-client.yaml managed-kubernetes.swagger.yaml
cd ..

if [ -n "$(git diff --stat)" ]; then
    echo "oapi generated code is outdated!"
    exit 1
fi

exit 0
