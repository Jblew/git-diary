package functions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jblew/git-diary/functions/util"
)

func verify(w http.ResponseWriter, r *http.Request) error {
	userData, err := util.GetUserAuthDataFromRequest(r)
	if err != nil {
		return err
	}
	if !application.IsUserAllowed(userData.Email) {
		return fmt.Errorf("User '%s' is not alowed", userData.Email)
	}

	log.Printf("User verified as %s\n", userData.Email)

	return nil
}
