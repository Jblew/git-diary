package app

import (
	"context"

	firebase "firebase.google.com/go"
)

// App is main class for functions
type App struct {
	Firebase *firebase.App
	Context  context.Context
}
