package gitpusher

// New initializes the GitPusher with given config
func New(config Config) (*GitPusher, error) {

	return &GitPusher{
		Config: config,
	}, nil
}
