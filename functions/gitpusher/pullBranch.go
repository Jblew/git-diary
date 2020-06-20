package gitpusher

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func (gitpusher *GitPusher) pullBranch(branchName string) error {
	err := gitpusher.worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branchName),
		Keep:   false,
	})
	if err != nil {
		return err
	}

	err = gitpusher.worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Force:      true,
	})

	if err != nil {
		return err
	}
	return nil
}
