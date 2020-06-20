package gitpusher

import (
	"github.com/go-git/go-git/v5"
)

// GitPusher allows pushing to a git repository
type GitPusher struct {
	Config   Config
	repo     *git.Repository
	worktree *git.Worktree
}
