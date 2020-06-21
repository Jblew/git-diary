#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

source ./functions.config.sh

if [ -z "${GCP_PROJECT_ID}" ]; then
    echo "GCP_PROJECT_ID env is not set"
    exit 1
fi

if [ -z "${ALLOWED_FIREBASE_USER_EMAILS}" ]; then
    echo "ALLOWED_FIREBASE_USER_EMAILS env is not set"
    exit 1
fi

if [ -z "${DIARY_REPOSITORY_URL}" ]; then
    echo "DIARY_REPOSITORY_URL env is not set"
    exit 1
fi

if [ -z "${BRANCH_NAME}" ]; then
    echo "BRANCH_NAME env is not set"
    exit 1
fi

if [ -z "${COMMIT_NAME}" ]; then
    echo "COMMIT_NAME env is not set"
    exit 1
fi

if [ -z "${COMMIT_EMAIL}" ]; then
    echo "COMMIT_EMAIL env is not set"
    exit 1
fi

if [ -z "${COMMIT_MESSAGE_FORMAT}" ]; then
    echo "COMMIT_MESSAGE_FORMAT env is not set"
    exit 1
fi

if [ -z "${DIARY_FILE_PATH_FORMAT}" ]; then
    echo "DIARY_FILE_PATH_FORMAT env is not set"
    exit 1
fi

if [ -z "${GIT_BASIC_AUTH_USERNAME}" ]; then
    echo "GIT_BASIC_AUTH_USERNAME env is not set"
    exit 1
fi

if [ -z "${GIT_BASIC_AUTH_PASSWORD}" ]; then
    echo "GIT_BASIC_AUTH_PASSWORD env is not set"
    exit 1
fi


GO_CONFIG_FILE="${DIR}/GetAutogeneratedConfig.go"

cat >"${GO_CONFIG_FILE}" <<EOF
package functions

import "github.com/Jblew/git-diary/functions/app"

// GetAutogeneratedConfig is automatically generated file with device config
func GetAutogeneratedConfig() app.Config {
	var conf = app.Config{}
	conf.ProjectID = "${GCP_PROJECT_ID}"
	conf.AllowedUserEmails = "${ALLOWED_FIREBASE_USER_EMAILS}"
    conf.RepositoryURL = "${DIARY_REPOSITORY_URL}"
    conf.BranchName = "${BRANCH_NAME}"
    conf.CommitName = "${COMMIT_NAME}"
    conf.CommitEmail = "${COMMIT_EMAIL}"
    conf.CommitMessageFormat = "${COMMIT_MESSAGE_FORMAT}"
    conf.DiaryFilePathFormat = "${DIARY_FILE_PATH_FORMAT}"
    conf.GitBasicAuthUsername = "${GIT_BASIC_AUTH_USERNAME}"
    conf.GitBasicAuthPassword = "${GIT_BASIC_AUTH_PASSWORD}"

	return conf
}

EOF
