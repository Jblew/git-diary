#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

source ./functions.config.sh

if [ -z "${GCP_PROJECT_ID}" ]; then
    echo "GCP_PROJECT_ID env is not set"
    exit 1
fi

if [ -z "${GCP_PROJECT_REGION}" ]; then
    echo "GCP_PROJECT_REGION env is not set"
    exit 1
fi


./generate-config.sh

gcloud config set project ${GCP_PROJECT_ID}

gcloud functions deploy PublishEntry \
  --trigger-http \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "1024MB"

