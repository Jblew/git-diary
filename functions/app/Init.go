package app

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/Jblew/git-diary/functions/gitpusher"
)

// Init initializes the app
func Init(config Config) (*App, error) {
	ctx := context.Background()
	conf := &firebase.Config{}

	firebase, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("Cannot initialize firebase.NewApp: %v", err)
	}

	firebaseAuth, err := firebase.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot initialize firebase.Auth: %v", err)
	}

	pusherConfig := gitpusher.Config{
		RepositoryURL: config.RepositoryURL,
	}
	gitPusher, err := gitpusher.New(pusherConfig)
	if err != nil {
		return nil, fmt.Errorf("Cannot initialize GitPusher: %v", err)
	}

	return &App{
		Firebase:     firebase,
		FirebaseAuth: firebaseAuth,
		Context:      ctx,
		Config:       config,
		GitPusher:    gitPusher,
	}, nil
}
