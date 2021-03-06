package functions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jblew/git-diary/functions/util"
)

// ReadDiary reads the diary file
func ReadDiary(writer http.ResponseWriter, req *http.Request) {
	out, err := handleReadDiary(writer, req)
	if err != nil {
		writer.WriteHeader(500)
		util.SendJSONError(writer, fmt.Sprintf("%v", err))
		return
	}

	fmt.Fprintln(writer, out)
}

func handleReadDiary(w http.ResponseWriter, r *http.Request) (string, error) {
	log.Println(util.DebugHTTPRequest(r))

	err := verify(w, r)
	if err != nil {
		return "", err
	}

	out, err := readDiaryFileFromRepo(w, r)
	if err != nil {
		return "", err
	}

	return out, nil
}

func readDiaryFileFromRepo(w http.ResponseWriter, r *http.Request) (string, error) {
	application.SyncLock()
	defer application.SyncUnlock()

	err := application.GitPusher.FetchMain()
	if err != nil {
		return "", fmt.Errorf("Cannot fetch main branch: %v", err)
	}

	diaryFilePath := util.FormatWithDate(application.Config.DiaryFilePathFormat)
	return application.GitPusher.ReadFile(diaryFilePath)
}
