package routineCommands

import (
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/channelExtends/protocol"

	"gitee.com/andyxt/gona/logger"
)

type ServerRoutineInboundCmdMsgRecv struct {
	routinePoolId int64
	routineId     int64
	Data          protocol.IProtocol
	Ctx           extends.OutterChannelHandlerContext
}

func NewServerRoutineInboundCmdMsgRecv(routinePoolId int64, routineId int64, Data protocol.IProtocol,
	Ctx extends.OutterChannelHandlerContext) (this *ServerRoutineInboundCmdMsgRecv) {
	this = new(ServerRoutineInboundCmdMsgRecv)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.Data = Data
	this.Ctx = Ctx
	return this
}
func (this *ServerRoutineInboundCmdMsgRecv) PoolId() int64 {
	return this.routinePoolId
}
func (this *ServerRoutineInboundCmdMsgRecv) QueueId() int64 {
	return this.routineId
}
func (this *ServerRoutineInboundCmdMsgRecv) Wait() (interface{}, bool) {
	return nil, true
}
func (this *ServerRoutineInboundCmdMsgRecv) Exec() {
	logger.Debug("ServerRoutineInboundCmdMsgRecv Exec")
}
