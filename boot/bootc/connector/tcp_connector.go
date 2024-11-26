package connector

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

type SocketType int

const (
	NormalSocket SocketType = 0
	WebSocket    SocketType = 1
)

type IConnector interface {
	Connect(ip string, port int, success IConnectSuccess, fail IConnectFail)
}

func NewTcpConnector(socketType SocketType, routinePoolID int64) IConnector {
	switch socketType {
	case WebSocket:
		instance := new(websocketConnector)
		instance.routinePoolID = routinePoolID
		return instance
	}
	instance := new(tcpConnector)
	instance.routinePoolID = routinePoolID
	return instance
}

type tcpConnector struct {
	routinePoolID int64
}

func (connector *tcpConnector) Connect(ip string, port int, success IConnectSuccess, fail IConnectFail) {
	routinePool.FireEvent(newConnectEvent(connector.routinePoolID, ip, port, -1, success, fail))
}

type websocketConnector struct {
	routinePoolID int64
}

func (connector *websocketConnector) Connect(ip string, port int, success IConnectSuccess, fail IConnectFail) {
	routinePool.FireEvent(newWebsocketConnectEvent(connector.routinePoolID, ip, port, -1, success, fail))
}

var routinePool *RoutinePool = NewRoutinePool(8, 16)

type RoutinePool struct {
	poolSize    int64
	chanSize    int64
	routinePool map[int64]*routine
}

func NewRoutinePool(PoolSize int64, ChanSize int64) (pool *RoutinePool) {
	pool = new(RoutinePool)
	pool.poolSize = PoolSize
	pool.chanSize = ChanSize
	pool.routinePool = make(map[int64]*routine)
	var routineId int64 = 0
	for routineId = 0; routineId < PoolSize; routineId = routineId + 1 {
		routine := NewRoutine(ChanSize)
		startChan := make(chan int, 1)
		routine.Start(startChan)
		<-startChan
		pool.routinePool[routineId] = routine
	}
	return
}
func (this *RoutinePool) ShutDown() {
	for _, value := range this.routinePool {
		value.Close()
	}
}
func (this *RoutinePool) FireEvent(e Event) {
	this.Exec(e)
}

/*发送内部STS消息(玩家不一定在线)*/
func (this *RoutinePool) FireEventWait(e Event) (result interface{}, ok bool) {
	this.Exec(e)
	return e.Wait()
}
func (this *RoutinePool) Exec(event Event) {
	routineId := event.QueueId() % this.poolSize
	if routine, ok := this.routinePool[routineId]; ok {
		routine.Put(event)
		return
	}
	logger.Error("event.QueueId()=", event.QueueId(), " % this.poolSize=", this.poolSize, " routineId==", routineId, " , but has no routine", " event:", event)
	if routine, ok := this.routinePool[0]; ok {
		routine.Put(event)
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

func (this *routine) Put(event Event) {
	this.queue <- event
}

func (this *routine) Start(startChan chan int) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	go func() {
		startChan <- 1
		for {
			event, ok := <-this.queue
			if event != nil {
				this.ExecEvent(event)
			}
			if !ok {
				break
			}
		}
	}()
}

func (this *routine) Close() {
	close(this.queue)
}

func (this *routine) ExecEvent(event Event) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	event.Exec()
}
