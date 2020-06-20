package util

import (
	"fmt"
	"net/http"
)

// UserAuthData holds authorization data for the logged in user
type UserAuthData struct {
	Email string
	UID   string
}

// GetUserAuthDataFromRequest extracts firebase user email from the http request.
// Returns UserAuthData if authenticated, error if not authenticated
func GetUserAuthDataFromRequest(r *http.Request) (UserAuthData, error) {
	userAuthData := UserAuthData{}
	userAuthData.UID = r.Header.Get("x-appengine-user-id")
	if userAuthData.UID == "" {
		return userAuthData, fmt.Errorf("Not authenticated")
	}
	userAuthData.Email = r.Header.Get("x-appengine-user-email")
	return userAuthData, nil
}

/*
Headers associated with GCP functions authorized request:

x-appengine-country: PL
x-appengine-region: mz
x-appengine-request-log-id: ...
x-appengine-user-nickname: ...
x-appengine-default-version-hostname: ...
x-appengine-user-id: ...
x-appengine-https: on
x-appengine-user-email: ...
function-execution-id: ...
x-appengine-city: ...
x-appengine-user-is-admin: 0
x-appengine-auth-domain: ...
x-appengine-timeout-ms: ...
x-cloud-trace-context: ...
x-forwarded-for: ...
sec-fetch-user: ...
x-appengine-citylatlong: ...
x-forwarded-proto: https
x-appengine-user-ip: ...
x-appengine-user-organization:
*/
