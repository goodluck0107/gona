package eventBus

import (
	"fmt"

	"gona/executor"

	"github.com/mohae/deepcopy"
)

type EventCallback func(...interface{})

var onEvents map[interface{}][]EventCallback = make(map[interface{}][]EventCallback) // call EventCallback after event trigged

// On is to register callback on events
// Only allowed to be executed during initialization
func On(ev string, f func(i ...interface{})) {
	// fmt.Println("eventBus On:", ev)
	onEvents[ev] = append(onEvents[ev], f)
}

// Trigger is to trigger an event with args
func Trigger(ev string, i ...interface{}) {
	// fmt.Println("eventBus Trigger:", ev)
	data := deepcopy.Copy(i).([]interface{})
	for _, f := range onEvents[ev] {
		copyData := deepcopy.Copy(data).([]interface{})
		f(copyData...)
	}
}

// TriggerCross is to trigger an event with args across player
func TriggerCross(ev string, poolId, uID int64, i ...interface{}) {
	// fmt.Println("eventBus TriggerCross:", ev)
	executor.FireEvent(newStsEvent(poolId, uID, false,
		ev, deepcopy.Copy(i).([]interface{})))
}

// TriggerCrossWait is to trigger an event with args across player and waiting for the result
func TriggerCrossWait(ev string, poolId, uID int64, i ...interface{}) {
	// fmt.Println("eventBus TriggerCrossWait:", ev)
	result, ok := executor.FireEventWait(newStsEvent(poolId, uID, true,
		ev, deepcopy.Copy(i).([]interface{})))
	if ok {
		fmt.Println("eventBus TriggerCrossWait result:", result)
	}
}
