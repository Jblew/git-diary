package gitpusher

import "github.com/go-git/go-billy/v5/memfs"

// New initializes the GitPusher with given config
func New(config Config) (*GitPusher, error) {

	fs := memfs.New()

	return &GitPusher{
		Config: config,
		fs:     fs,
	}, nil
}
