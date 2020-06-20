package gitpusher

import (
	"fmt"
	"log"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

// cloneIfNeeded clones repo and retrives the history
func (gitpusher *GitPusher) cloneIfNeeded() error {
	log.Println("Inside cloneIfNeeded")
	if gitpusher.repo == nil {
		log.Println("About to call doClone")
		repo, err := gitpusher.doClone()
		if err != nil {
			return err
		}
		log.Println("doClone done")

		log.Printf("Creating worktre, gitpusher: %+v", gitpusher)
		worktree, err := repo.Worktree()
		if err != nil {
			return fmt.Errorf("Could not initialize worktree: %v", err)
		}
		log.Println("Worktree created")

		gitpusher.repo = repo
		gitpusher.worktree = worktree
	}
	log.Println("cloneIfNeeded exitted if switch")
	return nil
}

func (gitpusher *GitPusher) doClone() (*git.Repository, error) {
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:

	fs := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
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
