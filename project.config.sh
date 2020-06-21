#!/usr/bin/env bash
PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

source "${PROJECT_DIR}/secrets.config.sh"

export GCP_PROJECT_ID="git-diary"
export GCP_PROJECT_REGION="europe-west1"
export GCP_FUNCTIONS_REGION="us-central1"
export DIARY_REPOSITORY_URL="https://github.com/doctor-lewandowski/ephemeris.git"
export BRANCH_NAME="master"
export DIARY_FILE_PATH_FORMAT="ephemeris-[yyyy]/[mm].md"
export COMMIT_NAME="Jędrzej Lewandowski"
export COMMIT_EMAIL="jedrzejblew@gmail.com"
export COMMIT_MESSAGE_FORMAT="Introitus ad ephemeridis [yyyy]/[mm]/[dd]"
export DIARY_ENTRY_CAPTION_FORMAT="\n> *[yyyy]/[mm]/[dd] [hh]:[ii]:[ss] UTC*\n\n"
