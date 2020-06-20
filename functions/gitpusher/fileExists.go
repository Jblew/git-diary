package gitpusher

import "os"

func (gitpusher *GitPusher) fileExists(path string) (bool, error) {
	_, err := gitpusher.fs.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
