package boost

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type VirtualTime interface {
	Now() time.Time
	Time(ts int64) time.Time
	Parse(string) (time.Time, error)
	Format(time.Time) string
	Location() *time.Location
	FakeDuration() time.Duration
}

const (
	datetimeParse  = "2006-1-2 15:04:05"
	dateParse      = "2006-1-2"
	datetimeFormat = "2006-01-02 15:04:05"
)

var virtualTime VirtualTime = NewFakeTime()

func SetVirtualTime(vt VirtualTime) {
	virtualTime = vt
}

type defaultVirtualTime struct{}

func (defaultVirtualTime) Parse(s string) (time.Time, error) {
	var format string
	if len(s) > 10 {
		format = datetimeParse
	} else {
		format = dateParse
	}
	return time.Parse(format, s)
}

func (defaultVirtualTime) Format(tm time.Time) string {
	return tm.Format(datetimeFormat)
}

func (defaultVirtualTime) Now() time.Time {
	return time.Now()
}

func (defaultVirtualTime) Time(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func (defaultVirtualTime) Location() *time.Location {
	return time.Local
}

func (defaultVirtualTime) FakeDuration() time.Duration {
	return 0
}

// FakeTime is a simulator for `FakeTime` in C language
type FakeTime struct {
	location     *time.Location
	fakeDuration time.Duration
}

func NewFakeTime() *FakeTime {
	v := &FakeTime{}
	v.initTimeLocation()
	v.initFakeTime()
	return v
}

func (v *FakeTime) initTimeLocation() {
	v.location = time.FixedZone("AWS", -4*3600)
}

func (v *FakeTime) initFakeTime() {
	s := os.Getenv("FAKETIME")
	if len(s) == 0 {
		return
	}
	if s[:1] == "-" || s[:1] == "+" {
		duration, err := strconv.ParseInt(s[:len(s)-1], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Environment variable FAKETIME is in wrong format: %s", s))
		}
		switch s[len(s)-1:] {
		case "s":
			v.fakeDuration = time.Duration(duration * 1e9)
		case "m":
			v.fakeDuration = time.Duration(duration * 1e9 * 60)
		case "h":
			v.fakeDuration = time.Duration(duration * 1e9 * 3600)
		case "d":
			v.fakeDuration = time.Duration(duration * 1e9 * 3600 * 24)
		case "w":
			v.fakeDuration = time.Duration(duration * 1e9 * 3600 * 24 * 14)
		default:
			panic(fmt.Sprintf("Environment variable FAKETIME is in wrong format: %s", s))
		}
	} else {
		var format string
		if len(s) > 10 {
			format = datetimeParse
		} else {
			format = dateParse
		}
		tm, err := time.ParseInLocation(format, s, v.location)
		if err != nil {
			panic(fmt.Sprintf("Environment variable FAKETIME is in wrong format: %s", s))
		}
		v.fakeDuration = time.Until(tm)
	}
	fmt.Printf("Environment variable FAKETIME is in correct format, faketime: %s, time is %s.\n", s, v.Format(v.Now()))
}

func (v *FakeTime) Time(ts int64) time.Time {
	return time.Unix(ts, 0).In(v.location)
}

func (v *FakeTime) Now() time.Time {
	return time.Now().Add(v.fakeDuration).In(v.location)
}

func (v *FakeTime) Parse(s string) (time.Time, error) {
	var format string
	if len(s) > 10 {
		format = datetimeParse
	} else {
		format = dateParse
	}
	return time.ParseInLocation(format, s, v.location)
}

func (v *FakeTime) Format(tm time.Time) string {
	return tm.Format(datetimeFormat)
}

func (v *FakeTime) Location() *time.Location {
	return v.location
}

func (v *FakeTime) FakeDuration() time.Duration {
	return v.fakeDuration
}

// SpecificTimestamp use specific timestamp
type SpecificTimestamp struct {
	FakeTime
	now time.Time
}

func NewSpecificTimestamp(tm int64) *SpecificTimestamp {
	v := &SpecificTimestamp{}
	v.initTimeLocation()
	v.now = time.Unix(tm, 0)
	fmt.Printf("Specific timestamp is set, timestamp: %v, time is %s.\n", tm, v.Format(v.Now()))
	return v
}

func (v *SpecificTimestamp) Now() time.Time {
	return v.now.In(v.location)
}

func (v *SpecificTimestamp) FakeDuration() time.Duration {
	return 0
}
