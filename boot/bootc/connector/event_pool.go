package connector

import (
	"fmt"

	"github.com/goodluck0107/gona/internal/logger"
	"github.com/goodluck0107/gona/utils"
)

type Event interface {
	/**
	 * 为保证事件序列化执行，需要序列化执行的事件必须提供一致的queueId
	 * */
	QueueId() (queueId int64)
	Exec()
}

var eventPool *routinePool = newRoutinePool(4, 8)

type routinePool struct {
	poolSize int64
	chanSize int64
	routines map[int64]*routine
}

func newRoutinePool(PoolSize int64, ChanSize int64) (pool *routinePool) {
	pool = new(routinePool)
	pool.poolSize = PoolSize
	pool.chanSize = ChanSize
	pool.routines = make(map[int64]*routine)
	var routineId int64 = 0
	for routineId = 0; routineId < PoolSize; routineId = routineId + 1 {
		routine := NewRoutine(ChanSize)
		startChan := make(chan int, 1)
		routine.Start(startChan)
		<-startChan
		pool.routines[routineId] = routine
	}
	return
}
func (pool *routinePool) ShutDown() {
	for _, value := range pool.routines {
		value.Close()
	}
}

func (pool *routinePool) FireEvent(e Event) {
	routineId := e.QueueId() % pool.poolSize
	if routine, ok := pool.routines[routineId]; ok {
		routine.Put(e)
		return
	}
	logger.Warn("event.QueueId()=", e.QueueId(), " poolSize=", pool.poolSize, " , but has no routine for ", " routineId=", routineId, " event:", e)
	if routine, ok := pool.routines[0]; ok {
		routine.Put(e)
	}
}

type routine struct {
	queue chan Event
}

func NewRoutine(ChanSize int64) (ret *routine) {
	ret = new(routine)
	ret.queue = make(chan Event, ChanSize)
	return
}

func (r *routine) Put(event Event) {
	r.queue <- event
}

func (r *routine) Start(startChan chan int) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	go func() {
		startChan <- 1
		for {
			event, ok := <-r.queue
			if event != nil {
				r.ExecEvent(event)
			}
			if !ok {
				break
			}
		}
	}()
}

func (r *routine) Close() {
	close(r.queue)
}

func (r *routine) ExecEvent(event Event) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	event.Exec()
}
