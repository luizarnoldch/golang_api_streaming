#!/bin/bash

FOLDERS=($(ls -d lambdas/*/))

test_lambda() {
  for folder in "${FOLDERS[@]}"; do
    (
      aws lambda invoke --function-name StreamProyect-${folder} --payload file://events/${folder}/input.json --cli-binary-format raw-in-base64-out ./events/${folder}/output.json
      echo -e "\n"
      cat ./events/${folder}/output.json
      echo -e "\n"
    )
  done
}

test_lambda