package gitpusher

import (
	"fmt"
	"log"
)

// FetchMain fetches branch specified in config
func (gitpusher *GitPusher) FetchMain() error {
	log.Println("About to call cloneIfNeeded")
	err := gitpusher.cloneIfNeeded()
	if err != nil {
		return err
	}
	log.Println("Executed cloneIfNeeded")

	log.Println("Pulling branch")
	err = gitpusher.pullBranch(gitpusher.Config.Branch)
	if err != nil {
		return fmt.Errorf("Could not pull the branch: %v", err)
	}

	return nil
}
