package gitpusher

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func (gitpusher *GitPusher) fetchBranch(branchName string) error {
	refSpec := config.RefSpec(fmt.Sprintf("+refs/heads/main:+refs/remotes/origin/%s", branchName))
	err := gitpusher.repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Force:      true,
		RefSpecs:   []config.RefSpec{refSpec},
	})
	if err != nil {
		return err
	}
	return nil
}
