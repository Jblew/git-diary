package functions

import (
	"fmt"
	"net/http"

	"github.com/Jblew/git-diary/functions/util"
)

// ResetLocalRepo resets the state of local repo
func ResetLocalRepo(writer http.ResponseWriter, req *http.Request) {
	out, err := handleResetLocalRepo(writer, req)
	if err != nil {
		util.SendJSONError(writer, fmt.Sprintf("%v", err))
		return
	}

	fmt.Fprintln(writer, out)
}

func handleResetLocalRepo(w http.ResponseWriter, r *http.Request) (string, error) {
	err := verify(w, r)
	if err != nil {
		return "", err
	}

	err = application.ResetPusher()

	if err != nil {
		return "", err
	}

	return "Done", nil
}
