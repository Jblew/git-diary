#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

nvm use v12
npm install -g firebase-tools
npm --prefix="public" install
npm --prefix="public" run build
firebase deploy --only hosting
