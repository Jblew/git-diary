package app

// Config is functions main config
type Config struct {
	ProjectID         string
	AllowedUserEmails string
	RepositoryURL     string
	BranchName        string
	CommitName        string
	CommitEmail       string
	CommitMessage     string
	DiaryFilePath     string
}
