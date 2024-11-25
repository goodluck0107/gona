package routineCommands

import (
	"gona/channelExtends/extends"
	"gona/channelExtends/protocol"

	"gona/logger"
)

type ClientRoutineOutboundCmdMsgSend struct {
	routinePoolId int64
	routineId     int64
	Data          protocol.IProtocol
	OnClose       bool // 是否在消息发送完毕后关闭连接
	PoolKey       int64
	ChlCtx        extends.OutterChannelHandlerContext
	Desc          string
}

func NewClientRoutineOutboundCmdMsgSend(routinePoolId int64, routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) (this *ClientRoutineOutboundCmdMsgSend) {
	this = new(ClientRoutineOutboundCmdMsgSend)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.Data = Data
	this.OnClose = OnClose
	this.PoolKey = PoolKey
	this.ChlCtx = ChlCtx
	this.Desc = Desc
	return
}
func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) PoolId() int64 {
	return msgSendEvent.routinePoolId
}
func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) QueueId() int64 {
	return msgSendEvent.routineId
}
func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) Exec() {
	if msgSendEvent.ChlCtx == nil {
		logger.Error("ClientRoutineOutboundCmdMsgSend Exec", "Fail:", " ChlCtx == nil")
		return
	}
	if msgSendEvent.Data == nil {
		logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), "Fail:", "Data == nil")
		return
	}
	if extends.IsClose(msgSendEvent.ChlCtx) {
		logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), "Fail:", "extends.IsClose(ChlCtx)")
		return
	}
	logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), " Success")
	msgSendEvent.ChlCtx.Write(msgSendEvent.Data)
	if msgSendEvent.OnClose {
		logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), " Success & CloseChannel")
		extends.Close(msgSendEvent.ChlCtx)
	}
}
