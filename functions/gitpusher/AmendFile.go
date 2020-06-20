package gitpusher

//AmendFileParams — params for AmendFile function
type AmendFileParams struct {
	Path          string
	Amendment     string
	CommitMessage string
}

// AmendFile appends to file and commits the changes
func (gitpusher *GitPusher) AmendFile(params AmendFileParams) error {
	err := gitpusher.appendToFile(params.Path, params.Amendment)
	if err != nil {
		return err
	}

	err = gitpusher.commitFile(params.Path, params.CommitMessage)
	if err != nil {
		return err
	}

	err = gitpusher.push()
	if err != nil {
		return err
	}

	return nil
}
