package gitpusher

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// AmendFile appends to file and commits the changes
func (gitpusher *GitPusher) commitFile(path string, commitMessage string) error {
	err := gitpusher.appendToFile(path, commitMessage)
	if err != nil {
		return err
	}

	_, err = gitpusher.worktree.Add(path)
	if err != nil {
		return err
	}

	_, err = gitpusher.worktree.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  gitpusher.Config.CommitName,
			Email: gitpusher.Config.CommitEmail,
			When:  time.Now(),
		},
	})
	return nil
}
