package timex

import (
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
