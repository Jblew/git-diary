package gitpusher

import (
	"io/ioutil"
)

// ReadFile reads current file contents
func (gitpusher *GitPusher) ReadFile(path string) (string, error) {
	err := gitpusher.createFileIfNotExists(path)
	if err != nil {
		return "", err
	}

	file, err := gitpusher.fs.Open(path) // Default is READONLY
	if err != nil {
		return "", err
	}

	binData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	return string(binData), nil
}
