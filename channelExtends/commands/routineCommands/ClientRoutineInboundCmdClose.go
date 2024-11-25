package routineCommands

import (
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/session"

	"gitee.com/andyxt/gona/logger"
)

type ClientRoutineInboundCmdClose struct {
	routinePoolId int64
	routineId     int64
	uID           int64
}

func NewClientRoutineInboundCmdClose(routinePoolId int64, routineId int64, uID int64, Desc string) (this *ClientRoutineInboundCmdClose) {
	this = new(ClientRoutineInboundCmdClose)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.uID = uID
	return
}
func (closeEvent *ClientRoutineInboundCmdClose) PoolId() int64 {
	return closeEvent.routinePoolId
}
func (closeEvent *ClientRoutineInboundCmdClose) QueueId() int64 {
	return closeEvent.routineId
}
func (closeEvent *ClientRoutineInboundCmdClose) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (closeEvent *ClientRoutineInboundCmdClose) Exec() {
	logger.Debug("ClientRoutineInboundCmdClose Exec")
	//connectionKey := strconv.Itoa(int(id)) + ip + strconv.Itoa(port)
	iSession := session.GetSession(0, closeEvent.uID)
	if iSession == nil {
		logger.Debug("ClientRoutineInboundCmdClose iSession == nil")
		return
	}
	chlCtx := extends.GetChlCtx(iSession)
	if chlCtx != nil {
		logger.Debug("ClientRoutineInboundCmdClose Close")
		extends.Close(chlCtx)
	}
	session.RemoveSession(0, closeEvent.uID)
}
