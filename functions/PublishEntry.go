package functions

import (
	"fmt"
	"net/http"
)

// PublishEntry publishes entry to git repo
func PublishEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test")
	return nil
}
