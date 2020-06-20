#!/usr/bin/env bash
PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

source "${PROJECT_DIR}/secrets.config.sh"

export GCP_PROJECT_ID="git-diary"
export GCP_PROJECT_REGION="europe-west1"
export GCP_FUNCTIONS_REGION="us-central1"
export DIARY_REPOSITORY_URL="git@github.com:Jblew/git-diary.git"
