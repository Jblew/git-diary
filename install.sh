#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

echo "# Installing git-diary"
echo "# Deploying functions"
./functions/deploy.sh
echo "# Functions deploy done"

echo "# Deploying firestore"
./firestore/deploy.sh
echo "# Functions firestore done"

echo "# Deploying ui"
./ui/deploy.sh
echo "# Functions ui done"

echo "# Git-diary installation done"
