#!/bin/bash

export GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"

OPEN_API_FILE="https://raw.githubusercontent.com/ElasticEmail/elasticemail-go/master/api/openapi.yaml?raw=true"

openapi-generator-cli validate -i "${OPEN_API_FILE}"
openapi-generator-cli generate \
    -i "${OPEN_API_FILE}" \
    -g go \
    -p packageName=elasticemail \
    -p enumClassPrefix=true \
    --git-host github.com \
    --git-repo-id elasticemail-go \
    --git-user-id pourcheragh
