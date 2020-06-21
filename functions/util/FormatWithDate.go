package util

import (
	"strings"
	"time"
)

func FormatWithDate(format string) string {
	out := format
	usedTime := time.Now().UTC()
	out = strings.ReplaceAll(out, "[yyyy]", string(usedTime.Year()))
	out = strings.ReplaceAll(out, "[mm]", usedTime.Month().String())
	out = strings.ReplaceAll(out, "[dd]", string(usedTime.Day()))
	return out
}
