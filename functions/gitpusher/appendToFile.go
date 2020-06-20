package gitpusher

import (
	"fmt"
	"os"
)

func (gitpusher *GitPusher) appendToFile(path string, append string) error {
	file, err := gitpusher.worktree.Filesystem.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("Could not open file: %v", err)
	}

	_, err = file.Write([]byte(append))
	if err != nil {
		return fmt.Errorf("Could not write to file: %v", err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("Could not close the file: %v", err)
	}

	return nil
}
