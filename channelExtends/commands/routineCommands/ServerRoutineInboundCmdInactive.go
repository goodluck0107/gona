package routineCommands

import (
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/session"

	"gitee.com/andyxt/gona/logger"
)

type ServerRoutineInboundCmdInactive struct {
	routinePoolId int64
	routineId     int64
	ChlCtx        extends.OutterChannelHandlerContext
}

func NewServerRoutineInboundCmdInactive(routinePoolId int64, routineId int64, ChlCtx extends.OutterChannelHandlerContext) (this *ServerRoutineInboundCmdInactive) {
	this = new(ServerRoutineInboundCmdInactive)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	return
}
func (inactiveEvent *ServerRoutineInboundCmdInactive) PoolId() int64 {
	return inactiveEvent.routinePoolId
}
func (inactiveEvent *ServerRoutineInboundCmdInactive) QueueId() int64 {
	return inactiveEvent.routineId
}
func (inactiveEvent *ServerRoutineInboundCmdInactive) Wait() (interface{}, bool) {
	return nil, true
}
func (inactiveEvent *ServerRoutineInboundCmdInactive) Exec() {
	logger.Debug("ServerRoutineInboundCmdInactive Exec:", extends.ChannelContextToString(inactiveEvent.ChlCtx))
	if extends.IsConflict(inactiveEvent.ChlCtx) { // 已被挤下线
		logger.Debug(extends.UID(inactiveEvent.ChlCtx), "OnChannelInactive conflit")
		return
	}
	uId := extends.UID(inactiveEvent.ChlCtx)
	iSession := session.GetSession(0, uId)
	if iSession == nil {
		logger.Error("ServerRoutineInboundCmdInactive Exec:", "UID = ", uId, " OnChannelInactive:player == nil")
		return
	}
	//掉线
	session.RemoveSession(0, uId)
	logger.Debug("OnChannelInactive:offline success RemovePlayer UID:", uId)
}
