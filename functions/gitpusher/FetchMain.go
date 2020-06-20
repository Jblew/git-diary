package gitpusher

// FetchMain fetches branch specified in config
func (gitpusher *GitPusher) FetchMain() error {
	err := gitpusher.cloneIfNeeded()
	if err != nil {
		return err
	}

	err = gitpusher.fetchBranch(gitpusher.Config.Branch)
	if err != nil {
		return err
	}

	return nil
}
