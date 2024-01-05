#!/bin/bash

FOLDERS=($(ls -d lambdas/*/))

export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="0"

mkdir -p ../bin

build_lambda() {
    for folder in "${FOLDERS[@]}"; do
    (
        folder_name=$(basename "${folder}")
        cd "lambdas/$folder_name" || { echo "Failed to cd into lambdas/$folder_name"; exit 1; }
        go build -o bootstrap -tags lambda.norpc || { echo "Failed to build in lambdas/$folder_name"; exit 1; }
        zip ../../bin/${folder_name}.zip bootstrap || { echo "Failed to zip bootstrap in lambdas/$folder_name"; exit 1; }
        rm -rf bootstrap
    )
    done
}

build_lambda