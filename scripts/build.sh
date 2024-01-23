#!/bin/bash

FOLDERS_STREAMS=($(ls -d lambdas/streams/*/))

export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="0"

build_streams() {
    for folder in "${FOLDERS_STREAMS[@]}"; do
    (
        folder_name=$(basename "${folder}")
        cd "lambdas/streams/$folder_name" || { echo "Failed to cd into lambdas/$folder_name"; exit 1; }
        go build -o bootstrap -tags lambda.norpc || { echo "Failed to build in lambdas/$folder_name"; exit 1; }
        zip ../../../bin/${folder_name}.zip bootstrap || { echo "Failed to zip bootstrap in lambdas/$folder_name"; exit 1; }
        rm -rf bootstrap
    )
    done
}

build_streams