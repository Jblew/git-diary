#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

echo "# Installing git-diary"
echo "# Deploying functions"
./functions/deploy.sh
echo "# Functions deploy done"

echo "# Deploying ui"
./ui/deploy.sh
echo "# Functions ui done"

echo "# Git-diary installation done"
