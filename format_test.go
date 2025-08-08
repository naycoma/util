package util_test

import (
	"testing"
	"time"

	"github.com/naycoma/util"
	"github.com/stretchr/testify/assert"
)

func TestStrFTime(t *testing.T) {
	a := assert.New(t)

	// Test basic parsing
	expectedTime := time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)
	actualTime := util.StrFTime("2006-01-02 15:04:05")
	a.Equal(expectedTime, actualTime, "Basic parsing")

	// Test with different time
	expectedTime = time.Date(2023, 10, 26, 10, 30, 0, 0, time.Local)
	actualTime = util.StrFTime("2023-10-26 10:30:00")
	a.Equal(expectedTime, actualTime, "Different time")
}

func TestFormatDuration(t *testing.T) {
	a := assert.New(t)

	// Test positive duration with milliseconds
	d := 1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Millisecond
	a.Equal("01:02:03.456", util.FormatDuration(d), "Positive duration with milliseconds")

	// Test positive duration without milliseconds
	d = 10*time.Hour + 20*time.Minute + 30*time.Second
	a.Equal("10:20:30", util.FormatDuration(d), "Positive duration without milliseconds")

	// Test negative duration
	d = -1*time.Minute
	a.Equal("-00:01:00", util.FormatDuration(d), "Negative duration")

	// Test zero duration
	d = 0
	a.Equal("00:00:00", util.FormatDuration(d), "Zero duration")

	// Test duration with only milliseconds
	d = 123 * time.Millisecond
	a.Equal("00:00:00.123", util.FormatDuration(d), "Only milliseconds")

	// Test duration with more than 24 hours
	d = 25*time.Hour + 30*time.Minute
	a.Equal("25:30:00", util.FormatDuration(d), "More than 24 hours")
}
