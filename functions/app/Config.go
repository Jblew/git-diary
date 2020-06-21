package app

// Config is functions main config
type Config struct {
	ProjectID               string
	AllowedUserEmails       string
	RepositoryURL           string
	GitBasicAuthUsername    string
	GitBasicAuthPassword    string
	BranchName              string
	CommitName              string
	CommitEmail             string
	CommitMessageFormat     string
	DiaryFilePathFormat     string
	DiaryEntryCaptionFormat string
}
