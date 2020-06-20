package app

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
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

	app := &App{
		Firebase:     firebase,
		FirebaseAuth: firebaseAuth,
		Context:      ctx,
		Config:       config,
	}

	err = app.ResetPusher()
	if err != nil {
		return nil, err
	}

	return app, nil
}
