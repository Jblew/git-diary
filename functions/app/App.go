package app

import (
	"context"
	"sync"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"github.com/Jblew/git-diary/functions/gitpusher"
)

// App is main class for functions
type App struct {
	Firebase     *firebase.App
	FirebaseAuth *firebaseAuth.Client
	Context      context.Context
	Config       Config
	GitPusher    *gitpusher.GitPusher
	Mutex        sync.Mutex
}
