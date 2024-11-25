package timex

import "time"

const (
	datetimeParse  = "2006-1-2 15:04:05"
	dateParse      = "2006-1-2"
	datetimeFormat = "2006-01-02 15:04:05"
)

type DefaultVirtualTime struct{}

func (DefaultVirtualTime) Parse(s string) (time.Time, error) {
	var format string
	if len(s) > 10 {
		format = datetimeParse
	} else {
		format = dateParse
	}
	return time.Parse(format, s)
}

func (DefaultVirtualTime) Format(tm time.Time) string {
	return tm.Format(datetimeFormat)
}

func (DefaultVirtualTime) Now() time.Time {
	return time.Now()
}

func (DefaultVirtualTime) Time(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func (DefaultVirtualTime) Location() *time.Location {
	return time.Local
}

func (DefaultVirtualTime) FakeDuration() time.Duration {
	return 0
}
