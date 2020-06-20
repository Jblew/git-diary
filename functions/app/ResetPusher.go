package app

import (
	"fmt"

	"github.com/Jblew/git-diary/functions/gitpusher"
)

// ResetPusher resets the git pusher
func (app *App) ResetPusher() error {
	config := app.Config
	pusherConfig := gitpusher.Config{
		RepositoryURL: config.RepositoryURL,
		Branch:        config.BranchName,
		CommitName:    config.CommitName,
		CommitEmail:   config.CommitEmail,
		CommitMessage: config.CommitMessage,
		AuthUsername:  config.GitBasicAuthUsername,
		AuthPassword:  config.GitBasicAuthPassword,
	}
	gitPusher, err := gitpusher.New(pusherConfig)
	if err != nil {
		return fmt.Errorf("Cannot initialize GitPusher: %v", err)
	}

	app.GitPusher = gitPusher
	return nil
}
