package gitpusher

import (
	"fmt"
	"os"
)

func (gitpusher *GitPusher) appendToFile(path string, append string) error {
	err := gitpusher.createFileIfNotExists(path)
	if err != nil {
		return err
	}

	file, err := gitpusher.worktree.Filesystem.OpenFile(path, os.O_WRONLY, os.ModeAppend.Perm())
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
