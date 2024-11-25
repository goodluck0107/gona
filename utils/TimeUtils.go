package utils

import (
	"time"
	"fmt"
)
var DateFormat string = "2006-01-02 15:04:05"
var BaseTime int64 = 1403366431146

func DaysSince20140622(timeMills int64) (day int32) {
	abs := timeMills - BaseTime
	return int32(abs * int64(time.Millisecond) / int64(time.Hour) / int64(24))
}
func GetMonthDay() string {
	timeNow :=time.Now()
	return fmt.Sprintf("%02d%02d", timeNow.Month(), timeNow.Day())
}

func GetTimePeriod(pTime time.Time)string{
	timeMinue :=pTime.Minute()
	if timeMinue>=30{
		//"2006-01-02 15:04:05"
		return pTime.Format("15")+":30"
	}
	return pTime.Format("15")+":00"
}