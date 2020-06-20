package util

import (
	"fmt"
	"net/http"
	"strings"
)

// DebugHTTPRequest formats http request data to string
// Source: https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func DebugHTTPRequest(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	return strings.Join(request, "\n")
}
