package routineCommands

import (
	"github.com/gox-studio/gona/channelExtends/extends"
	"github.com/gox-studio/gona/channelExtends/protocol"

	"github.com/gox-studio/gona/logger"
)

type ServerRoutineOutboundCmdMsgSend struct {
	routinePoolId int64
	routineId     int64
	Data          protocol.IProtocol
	OnClose       bool // 是否在消息发送完毕后关闭连接
	PoolKey       int64
	ChlCtx        extends.OutterChannelHandlerContext
	Desc          string
}

func NewServerRoutineOutboundCmdMsgSend(routinePoolId int64, routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) (this *ServerRoutineOutboundCmdMsgSend) {
	this = new(ServerRoutineOutboundCmdMsgSend)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.Data = Data
	this.OnClose = OnClose
	this.PoolKey = PoolKey
	this.ChlCtx = ChlCtx
	this.Desc = Desc
	return
}
func (this *ServerRoutineOutboundCmdMsgSend) PoolId() int64 {
	return this.routinePoolId
}
func (this *ServerRoutineOutboundCmdMsgSend) QueueId() int64 {
	return this.routineId
}
func (this *ServerRoutineOutboundCmdMsgSend) Wait() (interface{}, bool) {
	return nil, true
}
func (this *ServerRoutineOutboundCmdMsgSend) Exec() {
	//	logger.Debug(ctx.GetPoolKey(), " OnChannelMsgSend:", data)
	//fmt.Println("want OnChannelMsgSend msgId",  fmt.Sprintf("0x%04x", data.GetMessageId()))
	if this.ChlCtx == nil {
		logger.Error("ServerRoutineOutboundCmdMsgSend Exec", "Fail:", " ChlCtx == nil")
		return
	}
	if this.Data == nil {
		logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), "Fail:", "Data == nil")
		return
	}
	if extends.IsClose(this.ChlCtx) {
		logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), "Fail:", "extends.IsClose(ChlCtx)")
		return
	}
	this.ChlCtx.Write(this.Data)
	//fmt.Println("over OnChannelMsgSend msgId",  fmt.Sprintf("0x%04x", data.GetMessageId()))
	if this.OnClose {
		logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), " Success & CloseChannel")
		extends.Close(this.ChlCtx)
		return
	}
	logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), " Success")
}
