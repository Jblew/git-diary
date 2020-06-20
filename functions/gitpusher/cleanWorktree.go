package gitpusher

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (gitpusher *GitPusher) cleanWorktree() error {
	err := gitpusher.cleanWorktreeUnstaged()
	if err != nil {
		return err
	}

	err = gitpusher.cleanWorktreeHardReset()
	if err != nil {
		return err
	}

	return nil
}

func (gitpusher *GitPusher) cleanWorktreeUnstaged() error {
	err := gitpusher.worktree.Clean(&git.CleanOptions{
		Dir: true,
	})
	if err != nil {
		return fmt.Errorf("Cannot clean worktree: %v", err)
	}
	return nil
}

func (gitpusher *GitPusher) cleanWorktreeHardReset() error {
	headRef, err := gitpusher.repo.Head()
	if err != nil {
		return fmt.Errorf("Cannot get HEAD reference: %v", err)

	}
	err = gitpusher.worktree.Reset(&git.ResetOptions{
		Mode:   git.HardReset,
		Commit: headRef.Hash(),
	})
	if err != nil {
		return fmt.Errorf("Cannot reset worktree: %v", err)
	}
	return nil
}
