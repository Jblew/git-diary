package app

import (
	"context"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
)

// App is main class for functions
type App struct {
	Firebase     *firebase.App
	FirebaseAuth *firebaseAuth.Client
	Context      context.Context
	Config       Config
}
