package util

import (
	"fmt"
	"log"
	"net/http"
)

// SendJSONError sends JSON formatted error to http.ResponseWriter
func SendJSONError(w http.ResponseWriter, errMessage string) {
	_, err := fmt.Fprintf(w, "{\"error\":\"%s\"}", errMessage)
	if err != nil {
		log.Panicf("Cannot SendJSONError to http: %v", err)
	}
}
