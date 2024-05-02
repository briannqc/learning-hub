#!/bin/bash

OWNER=briannqc
REPO=learning-hub
WORKFLOW_ID=manually-triggered-workflow.yaml

curl -L \
    -X POST \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${GITHUB_TOKEN}" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/$OWNER/$REPO/actions/workflows/$WORKFLOW_ID/dispatches \
    -d '{
    "ref": "main",
    "inputs": {
        "string": "Brian NQC",
        "number": "2023",
        "option": "Option 1",
        "boolean": true,
        "environment": "prod"
      }
    }'
