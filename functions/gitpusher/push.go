package gitpusher

import (
	"github.com/go-git/go-git/v5"
)

// AmendFile appends to file and commits the changes
func (gitpusher *GitPusher) push() error {
	err := gitpusher.repo.Push(&git.PushOptions{})
	if err != nil {
		return err
	}
	return nil
}
