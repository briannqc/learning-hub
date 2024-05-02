#!/bin/bash

# https://docs.github.com/en/rest/repos/repos#create-a-repository-dispatch-event

OWNER=briannqc
REPO=learning-hub
curl -L \
    -X POST \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${GITHUB_TOKEN}" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/$OWNER/$REPO/dispatches \
    -d '{
        "event_type": "build",
        "client_payload": {
            "env": "production",
            "custom_field1": "Custom value 1"
        }
    }'
