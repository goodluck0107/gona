package routineCommands

import (
	"gona/channelExtends/extends"
	"gona/channelExtends/protocol"
)

type ClientRoutineInboundCmdMsgRecv struct {
	routinePoolId int64
	routineId     int64
	Data          protocol.IProtocol
	Ctx           extends.OutterChannelHandlerContext
}

func NewClientRoutineInboundCmdMsgRecv(routinePoolId int64, routineId int64, Data protocol.IProtocol,
	Ctx extends.OutterChannelHandlerContext) (this *ClientRoutineInboundCmdMsgRecv) {
	this = new(ClientRoutineInboundCmdMsgRecv)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.Data = Data
	this.Ctx = Ctx
	return this
}
func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) PoolId() int64 {
	return msgRecvEvent.routinePoolId
}
func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) QueueId() int64 {
	return msgRecvEvent.routineId
}
func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) Exec() {

}
