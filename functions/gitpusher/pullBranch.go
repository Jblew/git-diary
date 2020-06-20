package gitpusher

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func (gitpusher *GitPusher) pullBranch(branchName string) error {
	log.Println("About to call Checkout")

	err := gitpusher.worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branchName),
		Keep:   false,
	})
	if err != nil {
		return fmt.Errorf("Could not execute git worktree checkout: %v", err)
	}
	log.Println("Checkout done. Pulling")

	err = gitpusher.worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Force:      true,
	})
	log.Println("Pull done")

	if err != nil {
		return fmt.Errorf("Could not execute git worktree pull: %v", err)
	}
	return nil
}
