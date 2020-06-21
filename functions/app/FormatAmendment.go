package app

import "github.com/Jblew/git-diary/functions/util"

// FormatAmendment formats amendment with caption
func (app *App) FormatAmendment(amendment string) string {
	out := amendment

	out += util.FormatWithDate(app.Config.DiaryEntryCaptionFormat)

	return out
}
