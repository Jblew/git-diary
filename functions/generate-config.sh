#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

source ./functions.config.sh

if [ -z "${GCP_PROJECT_ID}" ]; then
    echo "GCP_PROJECT_ID env is not set"
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

	return conf
}

EOF
