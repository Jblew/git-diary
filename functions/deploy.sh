#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

source ./functions.config.sh

if [ -z "${GCP_PROJECT_ID}" ]; then
    echo "GCP_PROJECT_ID env is not set"
    exit 1
fi

if [ -z "${GCP_FUNCTIONS_REGION}" ]; then
    echo "GCP_FUNCTIONS_REGION env is not set"
    exit 1
fi


# Populate config
./generate-config.sh
# Test the build
go build *.go

gcloud config set project ${GCP_PROJECT_ID}
gcloud services enable cloudfunctions.googleapis.com

# Deploy functions in parallel
gcloud functions deploy PublishEntry \
  --allow-unauthenticated \
  --trigger-http \
  --region "${GCP_FUNCTIONS_REGION}" \
  --runtime go113 \
  --max-instances 1 \
  --memory "2048MB" &

gcloud functions deploy ReadDiary \
  --allow-unauthenticated \
  --trigger-http \
  --region "${GCP_FUNCTIONS_REGION}" \
  --runtime go113 \
  --memory "2048MB" &

gcloud functions deploy ResetLocalRepo \
  --allow-unauthenticated \
  --trigger-http \
  --region "${GCP_FUNCTIONS_REGION}" \
  --runtime go113 \
  --max-instances 1 \
  --memory "2048MB" &
wait


