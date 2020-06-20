package gitpusher

import (
	"os"
)

func (gitpusher *GitPusher) appendToFile(path string, append string) error {
	err := gitpusher.createFileIfNotExists(path)
	if err != nil {
		return err
	}

	file, err := gitpusher.fs.OpenFile(path, os.O_WRONLY, os.ModeAppend.Perm())
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(append))
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
