package utils

import (
	"fmt"
	"log"
	"time"
)

func StartTimer(secondsCount int, timeHandler TimeHandler) {
	go func() {
		for {
			timerSleep((time.Duration(secondsCount)) * time.Second)
			fireTimeEvent(timeHandler)
		}
	}()
}

func timerSleep(d time.Duration) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprint("timerSleep panic", string(Stack(3))))
		}
	}()
	time.Sleep(d)
}

func fireTimeEvent(timeHandler TimeHandler) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprint("fireTimeEvent", string(Stack(3))))
		}
	}()
	timeHandler.OnTimeEvent()
}
