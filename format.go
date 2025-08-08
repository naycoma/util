package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
)

// StrFTime parses a string into a time.Time object.
// It is equivalent to SQL's `strftime('%s', ?, 'utc')` for parsing.
// It uses time.DateTime format and time.Local location.
//
// For example:
//
//	StrFTime("2006-01-02 15:04:05") = 2006-01-02 15:04:05 +0900 JST
func StrFTime(str string) time.Time {
	return lo.Must(time.ParseInLocation(time.DateTime, str, time.Local))
}

// FormatDuration formats a time.Duration into a string in HH:MM:SS.ms format.
// It handles negative durations and includes millisecond precision if present.
//
// For example:
//
//	FormatDuration(1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Millisecond) = "01:02:03.456"
//	FormatDuration(-1*time.Minute) = "-00:01:00"
func FormatDuration(d time.Duration) string {
	milliseconds := int(d.Milliseconds())
	isMinus := milliseconds < 0
	if isMinus {
		milliseconds = -milliseconds
	}
	seconds, ms := DivMod(milliseconds, 1000)
	minutes, ss := DivMod(seconds, 60)
	hours, mm := DivMod(minutes, 60)
	var buf strings.Builder
	if isMinus {
		buf.WriteString("-")
	}
	buf.WriteString(fmt.Sprintf("%02d:", hours))
	buf.WriteString(fmt.Sprintf("%02d:%02d", mm, ss))
	if ms > 0 {
		buf.WriteString(fmt.Sprintf(".%03d", ms))
	}
	return buf.String()
}