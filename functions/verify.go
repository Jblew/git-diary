package functions

import (
	"context"
	"fmt"
	"log"
	"net/http"

	functionsAuth "github.com/Jblew/go-firebase-auth-in-gcp-functions"
)

func verify(w http.ResponseWriter, r *http.Request) error {
	firebaseUser, err := functionsAuth.AuthenticateFirebaseUser(context.Background(), r, application.FirebaseAuth)
	if err != nil {
		return err
	}
	if !application.IsUserAllowed(firebaseUser.Email) {
		return fmt.Errorf("User '%s' is not alowed", firebaseUser.Email)
	}

	log.Printf("User verified as %s\n", firebaseUser.Email)

	return nil
}
