package util

import (
	"time"
)

func MinTime(t time.Time, u ...time.Time) time.Time {
	for _, v := range u {
		if v.Before(t) {
			t = v
		}
	}
	return t
}

func MaxTime(t time.Time, u ...time.Time) time.Time {
	for _, v := range u {
		if v.After(t) {
			t = v
		}
	}
	return t
}

// StartOfDay returns the start of the day (00:00:00) for the given time in its timezone.
//
// For example:
//
//	StartOfDay(2025-08-08 15:30:00 JST) = 2025-08-08 00:00:00 JST
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Clock returns the duration since the start of the day for the given time in its timezone, including sub-second precision.
//
// For example:
//
//	Clock(2025-08-08 15:30:00.123 JST) = 15h30m0.123s
func Clock(t time.Time) time.Duration {
	return t.Sub(StartOfDay(t))
}

// DayOfWeekInWeek returns the 0-indexed day of the week, where startOfWeek is 0.
//
// For example:
//
//	DayOfWeekInWeek(2025-01-08 (Wednesday), Monday) = 2
func DayOfWeekInWeek(t time.Time, startOfWeek time.Weekday) int {
	return PositiveMod(int(t.Weekday()) - int(startOfWeek), 7)
}

// Deprecated: Use time.Time.ISOWeek() instead.
// WeekOfYearISO returns the ISO 8601 week number of the year.
// Week 1 is the first week with at least 4 days in the new year, and starts on Monday.
func WeekOfYearISO(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// WeekOfMonth returns the week number of the month for the given time, based on a custom start day of the week.
// Week 1 starts with the first occurrence of `startOfWeek` in the month.
//
// For example:
//
//	WeekOfMonth(2025-08-09 (Saturday), Monday) = 2
func WeekOfMonth(t time.Time, startOfWeek time.Weekday) int {
	monthStart := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	// Get the start of the week containing the 1st of the month, based on startOfWeek
	weekStartOfMonth1 := getWeekStartDate(monthStart, startOfWeek)

	// Get the start of the week containing t, based on startOfWeek
	weekStartOfT := getWeekStartDate(t, startOfWeek)

	// Calculate the number of weeks from weekStartOfMonth1 to weekStartOfT
	// This assumes weekStartOfMonth1 is the start of Week 1.
			return DaysBetween(weekStartOfMonth1, weekStartOfT)/7 + 1
}

// getWeekStartDate returns the date of the 'startOfWeek' for the week containing 't'.
func getWeekStartDate(t time.Time, startOfWeek time.Weekday) time.Time {
	daysToSubtract := PositiveMod(int(t.Weekday()) - int(startOfWeek), 7)
	return t.AddDate(0, 0, -daysToSubtract)
}

// DaysBetween returns the number of days between two time.Time values, ignoring time components and considering only dates.
// It calculates the number of days from 'from' to 'to'.
//
// For example:
//
//	DaysBetween(2025-01-01, 2025-01-03) = 2
//	DaysBetween(2025-01-01, 2025-01-01) = 0
//	DaysBetween(2025-01-03, 2025-01-01) = -2
func DaysBetween(from, to time.Time) int {
	return int(StartOfDay(to).Sub(StartOfDay(from)).Hours() / 24)
}

// TruncateLocal truncates a time to a multiple of d in its local timezone.
//
// For example:
//
//	TruncateLocal(2025-08-08 15:30:00 JST, 24 * time.Hour) = 2025-08-08 00:00:00 JST
func TruncateLocal(t time.Time, d time.Duration) time.Time {
	_, offsetSec := t.Zone()
	offset := time.Duration(offsetSec) * time.Second
	return t.Truncate(d).Add(-offset)
}

// WithinRange checks if t is within ±tolerance (inclusive) of a time rounded to the nearest interval.
//
// For example:
//
//	WithinRange(t, 1h, 5m) checks if t is between 9:55:00 and 10:05:00 (inclusive) for a round hour.
func WithinRange(t time.Time, interval, tolerance time.Duration) bool {
	diff := t.Sub(t.Truncate(interval))
	return diff <= tolerance || diff >= interval-tolerance
}

func TimeRange(times []time.Time) time.Duration {
	if len(times) == 0 {
		return 0
	}
	min := times[0]
	max := times[0]
	for _, t := range times {
		if t.Before(min) {
			min = t
		}
		if t.After(max) {
			max = t
		}
	}
	return max.Sub(min)
}
