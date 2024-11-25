package timex

import (
	"github.com/gox-studio/gona/cast"

	"github.com/tidwall/gjson"
)

type Timex struct {
	Zone  int64
	Fake  int64
	Delta int64
	VirtualTime
}

var options = &Options{}
var defaultVirtualTime = &DefaultVirtualTime{}
var timex = &Timex{
	VirtualTime: defaultVirtualTime,
}

func Init(s string) *Timex {
	if result := gjson.Get(s, "zone"); result.Type != gjson.Null {
		withTimeZone(result.String())
	} else {
		withTimeZone("Asia/Shanghai")
	}

	if result := gjson.Get(s, "fake"); result.Type != gjson.Null {
		withFakeTime(result.String())
	} else {
		withFakeTime("0")
	}

	if result := gjson.Get(s, "delta"); result.Type != gjson.Null {
		withDeltaTime(result.String())
	} else {
		withDeltaTime("0")
	}

	timex.VirtualTime = NewFakeTime()

	return timex
}

func withTimeZone(tz string) {
	timex.Zone = options.parseTimeZone(tz)
}
func withFakeTime(s string) {
	timex.Fake = options.parseFakeTime(s)
}
func withDeltaTime(s string) {
	timex.Delta = options.parseDeltaTime(s)
}

func (*Options) parseTimeZone(s string) int64 {
	tz, err := cast.ToTimeZoneE(s)
	if err != nil {
		panic(err)
	}
	return cast.ToInt64(tz)
}

func (*Options) parseFakeTime(s string) int64 {
	d, err := cast.ToDurationE(s)
	if err != nil {
		panic(err)
	}
	n, err := cast.ToInt64E(d)
	if err != nil {
		panic(err)
	}
	return n
}

func (*Options) parseDeltaTime(s string) int64 {
	d, err := cast.ToDurationE(s)
	if err != nil {
		panic(err)
	}
	n, err := cast.ToInt64E(d)
	if err != nil {
		panic(err)
	}
	return n
}

type Options struct{}
