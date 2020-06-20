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
	fmt.Fprintln(w, util.DebugHTTPRequest(r))
	// fmt.Println(w, application.Context)
	// application.FirebaseAuth.GetUser()
	// fmt.Fprintln(w, "Test")
}
