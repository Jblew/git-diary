package app

import "strings"

// IsUserAllowed checks if user is on allowed user list
func (app *App) IsUserAllowed(email string) bool {
	allowedEmails := strings.Split(app.Config.AllowedUserEmails, ",")
	for _, allowedEmail := range allowedEmails {
		if allowedEmail == email {
			return true
		}
	}
	return false
}
