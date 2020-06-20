package gitpusher

// Config is a GitPusher config
type Config struct {
	RepositoryURL string
	AuthUsername  string
	AuthPassword  string
	Branch        string
	CommitName    string
	CommitEmail   string
	CommitMessage string
}
