package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jblew/git-diary/functions/gitpusher"
	"github.com/Jblew/git-diary/functions/util"
)

// PublishEntry publishes entry to git repo
func PublishEntry(writer http.ResponseWriter, req *http.Request) {
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

	entryBin, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("Cannot read request body: %v", err)
	}

	entry := string(entryBin)
	if len(entry) == 0 {
		return "", fmt.Errorf("You must provide entry to publish as text in request body")
	}

	out, err := pushEntryToRepo(entry, w, r)
	if err != nil {
		return "", err
	}

	return out, nil
}

func pushEntryToRepo(entry string, w http.ResponseWriter, r *http.Request) (string, error) {
	application.SyncLock()
	defer application.SyncUnlock()

	err := application.GitPusher.FetchMain()
	if err != nil {
		return "", fmt.Errorf("Cannot fetch main branch: %v", err)
	}

	log.Println("Repository fetch successfull")

	diaryFilePath := util.FormatWithDate(application.Config.DiaryFilePathFormat)
	commitMessage := util.FormatWithDate(application.Config.CommitMessageFormat)
	amendmentFormatted := application.FormatAmendment(entry)

	err = application.GitPusher.AmendFile(gitpusher.AmendFileParams{
		Path:          diaryFilePath,
		Amendment:     amendmentFormatted,
		CommitMessage: commitMessage,
	})
	if err != nil {
		return "", fmt.Errorf("Cannot amend to branch: %v", err)
	}
	log.Println("Repository file amended")

	return application.GitPusher.ReadFile(diaryFilePath)
}
