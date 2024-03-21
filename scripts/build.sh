#!/bin/bash

FOLDERS_API=($(ls -d lambdas/api/*/))

export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="0"

build_api() {
    for folder in "${FOLDERS_API[@]}"; do
    (
        folder_name=$(basename "${folder}")
        cd "lambdas/api/$folder_name" || { echo "Failed to cd into lambdas/api/$folder_name"; exit 1; }
        go build -o bootstrap -tags lambda.norpc || { echo "Failed to build in lambdas/api/$folder_name"; exit 1; }
        zip ../../../bin/${folder_name}.zip bootstrap || { echo "Failed to zip bootstrap in lambdas/api/$folder_name"; exit 1; }
        rm -rf bootstrap
    )
    done
}

build_api

# FOLDERS_COGNITO=($(ls -d lambdas/cognito/*/))

# export GOOS="linux"
# export GOARCH="amd64"
# export CGO_ENABLED="0"

# build_cognito() {
#     for folder in "${FOLDERS_COGNITO[@]}"; do
#     (
#         folder_name=$(basename "${folder}")
#         cd "lambdas/streams/$folder_name" || { echo "Failed to cd into lambdas/$folder_name"; exit 1; }
#         go build -o bootstrap -tags lambda.norpc || { echo "Failed to build in lambdas/$folder_name"; exit 1; }
#         zip ../../../bin/${folder_name}.zip bootstrap || { echo "Failed to zip bootstrap in lambdas/$folder_name"; exit 1; }
#         rm -rf bootstrap
#     )
#     done
# }

# build_cognito