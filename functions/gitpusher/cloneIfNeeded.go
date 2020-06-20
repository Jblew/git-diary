package gitpusher

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

// cloneIfNeeded clones repo and retrives the history
func (gitpusher *GitPusher) cloneIfNeeded() error {
	if gitpusher.repo == nil {
		repo, err := gitpusher.doClone()
		if err != nil {
			return err
		}

		worktree, err := gitpusher.repo.Worktree()
		if err != nil {
			return fmt.Errorf("Could not initialize worktree: %v", err)
		}

		gitpusher.repo = repo
		gitpusher.worktree = worktree
	}
	return nil
}

func (gitpusher *GitPusher) doClone() (*git.Repository, error) {

	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:

	repo, err := git.Clone(memory.NewStorage(), gitpusher.fs, &git.CloneOptions{
		URL:  gitpusher.Config.RepositoryURL,
		Auth: gitpusher.getCloneAuth(),
	})
	if err != nil {
		return nil, fmt.Errorf("Cannot clone: %v", err)
	}
	return repo, nil
}

func (gitpusher *GitPusher) getCloneAuth() *http.BasicAuth {
	return &http.BasicAuth{
		Username: gitpusher.Config.AuthUsername,
		Password: gitpusher.Config.AuthPassword,
	}
}
