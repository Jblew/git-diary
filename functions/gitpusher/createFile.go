package gitpusher

func (gitpusher *GitPusher) createFileIfNotExists(path string) error {
	exists, err := gitpusher.fileExists(path)
	if err != nil {
		return err
	} else if exists == false {
		return gitpusher.createFile(path)
	}
	return nil
}

func (gitpusher *GitPusher) createFile(path string) error {
	_, err := gitpusher.worktree.Filesystem.Create(path)
	if err != nil {
		return err
	}
	return nil
}
