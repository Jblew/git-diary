package functions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jblew/git-diary/functions/gitpusher"
	"github.com/Jblew/git-diary/functions/util"
)

// PublishEntry publishes entry to git repo
func PublishEntry(writer http.ResponseWriter, req *http.Request) {
	// application.ResetPusher()

	out, err := handlePublishEntry(writer, req)
	if err != nil {
		util.SendJSONError(writer, fmt.Sprintf("%v", err))
		return
	}

	fmt.Fprintln(writer, out)
}

func handlePublishEntry(w http.ResponseWriter, r *http.Request) (string, error) {
	err := verify(w, r)
	if err != nil {
		return "", err
	}

	out, err := pushEntryToRepo(w, r)
	if err != nil {
		return "", err
	}

	return out, nil
}

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

func pushEntryToRepo(w http.ResponseWriter, r *http.Request) (string, error) {
	err := application.GitPusher.FetchMain()
	if err != nil {
		return "", fmt.Errorf("Cannot fetch main branch: %v", err)
	}

	log.Println("Repository fetch successfull")

	err = application.GitPusher.AmendFile(gitpusher.AmendFileParams{
		Path:          application.Config.DiaryFilePath,
		Amendment:     "\n\n## test",
		CommitMessage: application.Config.CommitMessage,
	})
	if err != nil {
		return "", fmt.Errorf("Cannot amend to branch: %v", err)
	}
	log.Println("Repository file amended")

	return application.GitPusher.ReadFile(application.Config.DiaryFilePath)
}
