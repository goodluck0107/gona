package boost

import (
	"time"
)

func Time(ts int64) time.Time {
	return virtualTime.Time(ts)
}

func Now() time.Time {
	return virtualTime.Now()
}

func ParseTime(s string) time.Time {
	tm, err := virtualTime.Parse(s)
	if err != nil {
		panic(err)
	}
	return tm
}

func FormatTime(tm time.Time) string {
	return virtualTime.Format(tm)
}

func TimeLocation() *time.Location {
	return virtualTime.Location()
}

func FakeDuration() time.Duration {
	return virtualTime.FakeDuration()
}

// IsSameDay returns true if timestamp ts1 & ts2 in same day
func IsSameDay(ts1 int64, ts2 int64) bool {
	dt1 := Time(ts1)
	dt2 := Time(ts2)

	year1, month1, day1 := dt1.Date()
	year2, month2, day2 := dt2.Date()

	return year1 == year2 && month1 == month2 && day1 == day2
}

// IsToday returns true if timestamp ts in today
func IsToday(ts int64) bool {
	dt := Time(ts)
	now := Now()

	year, month, day := dt.Date()
	yearNow, monthNow, dayNow := now.Date()
	return year == yearNow && month == monthNow && day == dayNow
}

// IsSameDay returns true if timestamp ts1 & ts2 in same day
func IsSameWeek(ts1 int64, ts2 int64) bool {
	dt1 := Time(ts1)
	dt2 := Time(ts2)

	year1, week1 := dt1.ISOWeek()
	year2, week2 := dt2.ISOWeek()

	return year1 == year2 && week1 == week2
}

// IsThisWeek returns true if timestamp t in this day
func IsThisWeek(ts int64) bool {
	dt := Time(ts)
	now := Now()

	year, week := dt.ISOWeek()
	yearNow, weekNow := now.ISOWeek()

	return year == yearNow && week == weekNow
}

// SecondsOfThisWeek gets seconds passed in this week
func SecondsOfThisWeek() int64 {
	now := Now()
	weekDay := now.Weekday()
	day := int64(0)
	if weekDay == time.Sunday {
		day = 6
	} else {
		day = int64(weekDay) - 1
	}
	return day*86400 + SecondsOfToday()
}

// SecondsOfToday get seconds passed of today
func SecondsOfToday() int64 {
	now := Now()
	return 60*60*int64(now.Hour()) + 60*int64(now.Minute()) + int64(now.Second())
}

// SecondsOfDay gets the beginning time of that day
func SecondsOfDay(ts int64) int64 {
	tm := Time(ts)
	return 60*60*int64(tm.Hour()) + 60*int64(tm.Minute()) + int64(tm.Second())
}

// IntervalDays gets interval day from the beginning of that day to now
func IntervalDays(ts int64) int64 {
	return (Now().Unix() - ts + SecondsOfDay(ts)) / 86400
}

func BeginningOfDay(ts int64) int64 {
	return ts - SecondsOfDay(ts)
}

func EndOfDay(ts int64) int64 {
	return BeginningOfDay(ts) + 3600*24
}
