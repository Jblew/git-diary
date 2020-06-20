package gitpusher

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

// CloneAndGetHistory clones repo and retrives the history
func (gitpusher *GitPusher) CloneAndGetHistory() (string, error) {
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: gitpusher.Config.RepositoryURL,
	})
	if err != nil {
		return "", fmt.Errorf("Cannot clone: %v", err)
	}

	headRef, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("Cannot get HEAD: %v", err)
	}

	logIterator, err := repo.Log(&git.LogOptions{From: headRef.Hash()})
	if err != nil {
		return "", fmt.Errorf("Cannot get log iterator: %v", err)
	}

	out := ""
	// ... just iterates over the commits, printing it
	err = logIterator.ForEach(func(c *object.Commit) error {
		out += c.Message
		out += "\n"
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("Cannot get git history: %v", err)
	}
	return out, nil
}
