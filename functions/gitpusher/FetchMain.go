package gitpusher

import "fmt"

// FetchMain fetches branch specified in config
func (gitpusher *GitPusher) FetchMain() error {
	err := gitpusher.cloneIfNeeded()
	if err != nil {
		return err
	}

	err = gitpusher.pullBranch(gitpusher.Config.Branch)
	if err != nil {
		return fmt.Errorf("Could not pull the branch: %v", err)
	}

	return nil
}
