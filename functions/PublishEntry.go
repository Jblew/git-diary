package functions

import (
	"fmt"
	"net/http"

	"github.com/Jblew/git-diary/functions/util"
)

// PublishEntry publishes entry to git repo
func PublishEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, util.DebugHTTPRequest(r))
	// application.FirebaseAuth.GetUser()
	// fmt.Fprintln(w, "Test")
}
