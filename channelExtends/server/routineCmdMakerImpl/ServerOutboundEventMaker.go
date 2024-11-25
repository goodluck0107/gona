package routineCmdMakerImpl

import (
	"gona/channelExtends/commands/routineCommands"
	"gona/channelExtends/extends"
	"gona/channelExtends/protocol"

	"gona/executor"
)

type ServerOutboundEventMaker struct {
}

func NewServerOutboundEventMaker() (impl *ServerOutboundEventMaker) {
	impl = new(ServerOutboundEventMaker)
	return
}

func (impl *ServerOutboundEventMaker) MakeConnectEvent(routinePoolId int64, routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event {
	return nil
}

func (impl *ServerOutboundEventMaker) MakeCloseEvent(routinePoolId int64, routineId int64, uID int64, Desc string) executor.Event {
	return nil
}

func (impl *ServerOutboundEventMaker) MakeSendMessageEvent(routinePoolId int64, routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) executor.Event {
	return routineCommands.NewServerRoutineOutboundCmdMsgSend(routinePoolId, routineId, Data, OnClose, PoolKey, ChlCtx, Desc)
}
