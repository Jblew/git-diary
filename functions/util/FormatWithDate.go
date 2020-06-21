package util

import (
	"fmt"
	"strings"
	"time"
)

func FormatWithDate(format string) string {
	out := format
	usedTime := time.Now().UTC()
	out = strings.ReplaceAll(out, "[yyyy]", fmt.Sprintf("%d", usedTime.Year()))
	out = strings.ReplaceAll(out, "[mm]", fmt.Sprintf("%02d", usedTime.Month()))
	out = strings.ReplaceAll(out, "[dd]", fmt.Sprintf("%02d", usedTime.Day()))
	return out
}
