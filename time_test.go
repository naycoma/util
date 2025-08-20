package util_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/naycoma/util"
)

// date is a helper function for creating time.Time objects in the "Asia/Tokyo" timezone.
func date(year int, month time.Month, day int) time.Time {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

// dateUTC is a helper function for creating time.Time objects in UTC.
func dateUTC(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func parse(value string) time.Time {
	return lo.Must(time.Parse(time.RFC3339, value))
}

func TestTruncateLocal(t *testing.T) {
	a := assert.New(t)
	now := parse("2006-01-02T15:04:05+09:00")
	_, offsetSec := now.Zone()
	offset := time.Duration(offsetSec) * time.Second

	t.Run("Truncate", func(t *testing.T) {
		floor := now.Truncate(time.Hour * 24)
		a.Equal(floor.Hour(), int(offset.Hours()))
		a.Equal(floor.Minute(), 0)
		a.Equal(floor.Second(), 0)
	})

	t.Run("TruncateLocal", func(t *testing.T) {
		floor := util.TruncateLocal(now, time.Hour*24)
		a.Equal(floor.Hour(), 0)
		a.Equal(floor.Minute(), 0)
		a.Equal(floor.Second(), 0)
	})

	t.Run("TruncateUTC", func(t *testing.T) {
		floor := now.UTC().Truncate(time.Hour * 24)
		a.Equal(floor.Hour(), 0)
		a.Equal(floor.Minute(), 0)
		a.Equal(floor.Second(), 0)
	})
}

func TestClock(t *testing.T) {
	a := assert.New(t)

	testTime := date(2025, 8, 9).Add(15*time.Hour + 30*time.Minute + 45*time.Second + 123*time.Millisecond + 456*time.Microsecond + 789*time.Nanosecond)
	expectedDurationStr := "15h30m45.123456789s"

	duration := util.Clock(testTime)
	a.Equal(expectedDurationStr, duration.String())
}

func TestDaysBetween(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, util.DaysBetween(date(2025, 1, 1).Add(10*time.Hour), date(2025, 1, 3).Add(15*time.Hour)), "2 days difference")
	a.Equal(0, util.DaysBetween(date(2025, 1, 1), date(2025, 1, 1).Add(23*time.Hour+59*time.Minute+59*time.Second)), "Same day difference")
	a.Equal(4, util.DaysBetween(date(2025, 1, 1), date(2025, 1, 5)), "4 days difference")

	// Test across year boundary
	a.Equal(217, util.DaysBetween(date(2024, 12, 30), date(2025, 8, 4)), "Days between 2024-12-30 and 2025-08-04")
	a.Equal(217, util.DaysBetween(date(2024, 12, 29), date(2025, 8, 3)), "Days between 2024-12-29 and 2025-08-03")

	// Test negative difference
	a.Equal(-2, util.DaysBetween(date(2025, 1, 3).Add(15*time.Hour), date(2025, 1, 1).Add(10*time.Hour)), "Negative days difference")
}

func TestDayOfWeekInWeek(t *testing.T) {
	a := assert.New(t)

	// Test with UTC dates
	// Monday is 0, Tuesday is 1, ..., Sunday is 6
	a.Equal(0, util.DayOfWeekInWeek(dateUTC(2025, 1, 6), time.Monday), "UTC: 2025-01-06 is Monday")
	a.Equal(1, util.DayOfWeekInWeek(dateUTC(2025, 1, 7), time.Monday), "UTC: 2025-01-07 is Tuesday")
	a.Equal(6, util.DayOfWeekInWeek(dateUTC(2025, 1, 5), time.Monday), "UTC: 2025-01-05 is Sunday")

	// Sunday is 0, Monday is 1, ..., Saturday is 6
	a.Equal(0, util.DayOfWeekInWeek(dateUTC(2025, 1, 5), time.Sunday), "UTC: 2025-01-05 is Sunday")
	a.Equal(1, util.DayOfWeekInWeek(dateUTC(2025, 1, 6), time.Sunday), "UTC: 2025-01-06 is Monday")
	a.Equal(6, util.DayOfWeekInWeek(dateUTC(2025, 1, 4), time.Sunday), "UTC: 2025-01-04 is Saturday")

	// Test with local dates (Asia/Tokyo)
	// Monday is 0, Tuesday is 1, ..., Sunday is 6
	a.Equal(0, util.DayOfWeekInWeek(date(2025, 1, 6), time.Monday), "Local: 2025-01-06 is Monday")
	a.Equal(1, util.DayOfWeekInWeek(date(2025, 1, 7), time.Monday), "Local: 2025-01-07 is Tuesday")
	a.Equal(6, util.DayOfWeekInWeek(date(2025, 1, 5), time.Monday), "Local: 2025-01-05 is Sunday")

	// Sunday is 0, Monday is 1, ..., Saturday is 6
	a.Equal(0, util.DayOfWeekInWeek(date(2025, 1, 5), time.Sunday), "Local: 2025-01-05 is Sunday")
	a.Equal(1, util.DayOfWeekInWeek(date(2025, 1, 6), time.Sunday), "Local: 2025-01-06 is Monday")
	a.Equal(6, util.DayOfWeekInWeek(date(2025, 1, 4), time.Sunday), "Local: 2025-01-04 is Saturday")
}

func TestWeekOfYearISO(t *testing.T) {
	a := assert.New(t)

	// 2025-01-01 (水) -> ISO Week 1 (2024-12-30 - 2025-01-05)
	a.Equal(1, util.WeekOfYearISO(date(2025, 1, 1)))
	// 2025-01-05 (日) -> ISO Week 1
	a.Equal(1, util.WeekOfYearISO(date(2025, 1, 5)))
	// 2025-01-06 (月) -> ISO Week 2
	a.Equal(2, util.WeekOfYearISO(date(2025, 1, 6)))
	// 2025-08-09 (土) -> ISO Week 32
	a.Equal(32, util.WeekOfYearISO(date(2025, 8, 9)))

	// 2024-12-30 (月) -> ISO Week 1 of 2025
	a.Equal(1, util.WeekOfYearISO(date(2024, 12, 30)))
	// 2024-12-29 (日) -> ISO Week 52 of 2024
	a.Equal(52, util.WeekOfYearISO(date(2024, 12, 29)))
}

func TestWeekOfMonth(t *testing.T) {
	a := assert.New(t)

	// 2025年8月1日 (金曜日)
	timeAug1 := date(2025, 8, 1)

	// 2025年8月9日 (土曜日)
	timeAug9 := date(2025, 8, 9)

	// 週の始まりが月曜日の場合
	// 2025-08-01 (金) -> 2025-07-28 (月) が週の始まり
	a.Equal(1, util.WeekOfMonth(timeAug1, time.Monday), "Aug 1, Monday start")
	// 2025-08-09 (土) -> 2025-08-04 (月) が週の始まり
	a.Equal(2, util.WeekOfMonth(timeAug9, time.Monday), "Aug 9, Monday start")

	// 週の始まりが日曜日の場合
	// 2025-08-01 (金) -> 2025-07-27 (日) が週の始まり
	a.Equal(1, util.WeekOfMonth(timeAug1, time.Sunday), "Aug 1, Sunday start")
	// 2025-08-09 (土) -> 2025-08-03 (日) が週の始まり
	a.Equal(2, util.WeekOfMonth(timeAug9, time.Sunday), "Aug 9, Sunday start")
}

func TestFuturePast(t *testing.T) {
	a := assert.New(t)

	futureTime := time.Now().Add(1 * time.Hour)
	pastTime := time.Now().Add(-1 * time.Hour)

	a.True(util.Future(futureTime))
	a.False(util.Future(pastTime))

	a.False(util.Past(futureTime))
	a.True(util.Past(pastTime))
}
