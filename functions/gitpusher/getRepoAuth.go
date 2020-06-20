package gitpusher

import (
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func (gitpusher *GitPusher) getRepoAuth() *http.BasicAuth {
	return &http.BasicAuth{
		Username: gitpusher.Config.AuthUsername,
		Password: gitpusher.Config.AuthPassword,
	}
}
