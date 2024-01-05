#!/bin/bash

FOLDERS=($(ls -d lambdas/*/))

test_lambda() {
  for folder in "${FOLDERS[@]}"; do
    (
      folder_name=$(basename "${folder}")
      aws lambda invoke --function-name StreamProyect-${folder_name} --payload file://events/${folder_name}/input.json --cli-binary-format raw-in-base64-out ./events/${folder_name}/output.json
      echo -e "\n"
      cat ./events/${folder_name}/output.json
      echo -e "\n"
    )
  done
}

test_lambda