package functions

import (
	"fmt"
	"net/http"

	"github.com/Jblew/git-diary/functions/util"
)

// PublishEntry publishes entry to git repo
func PublishEntry(w http.ResponseWriter, r *http.Request) {
	userData, err := util.GetUserAuthDataFromRequest(r)
	if err != nil {
		util.SendJSONError(w, fmt.Sprintf("%v", err))
		return
	}
	if !application.IsUserAllowed(userData.Email) {
		util.SendJSONError(w, fmt.Sprintf("User '%s' is not alowed", userData.Email))
		return
	}

	repoHistory, err := application.GitPusher.CloneAndGetHistory()
	if err != nil {
		util.SendJSONError(w, fmt.Sprintf("Git pusher error: %v", err))
		return
	}

	fmt.Fprintln(w, repoHistory)
	// fmt.Println(w, application.Context)
	// application.FirebaseAuth.GetUser()
	// fmt.Fprintln(w, "Test")
}
