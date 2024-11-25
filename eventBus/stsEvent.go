package eventBus

import (
	"github.com/mohae/deepcopy"
)

type stsEvent struct {
	routinePoolId int64
	routineId     int64
	waitChan      chan int8
	evt           string
	params        []interface{}
}

func newStsEvent(routinePoolId int64, routineId int64, syncWait bool,
	evt string, params []interface{}) (this *stsEvent) {
	this = new(stsEvent)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	if syncWait {
		this.waitChan = make(chan int8, 1)
	}
	this.evt = evt
	this.params = params
	return
}
func (stsEvent *stsEvent) PoolId() int64 {
	return stsEvent.routinePoolId
}
func (stsEvent *stsEvent) QueueId() (queueId int64) {
	return stsEvent.routineId
}

func (stsEvent *stsEvent) Wait() (interface{}, bool) {
	if stsEvent.waitChan != nil {
		<-stsEvent.waitChan
	}
	return nil, true
}

func (stsEvent *stsEvent) Exec() {
	for _, f := range onEvents[stsEvent.evt] {
		copyData := deepcopy.Copy(stsEvent.params).([]interface{})
		f(copyData...)
	}
	if stsEvent.waitChan != nil {
		stsEvent.waitChan <- 1
	}
}
