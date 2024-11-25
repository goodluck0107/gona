package routineCmdMakerImpl

import (
	"gona/channelExtends/commands/routineCommands"
	"gona/channelExtends/extends"
	"gona/channelExtends/protocol"

	"gona/bootStrap/bootStrapClient/listener"
	"gona/channel"
	"gona/executor"
)

type ClientOutboundEventMaker struct {
	connector listener.IConnector
}

func NewClientOutboundEventMaker(connector listener.IConnector) (impl *ClientOutboundEventMaker) {
	impl = new(ClientOutboundEventMaker)
	impl.connector = connector
	return
}

func (impl *ClientOutboundEventMaker) MakeConnectEvent(routinePoolId int64, routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdConnect(routinePoolId, routineId, uID, ip, port, params, impl.connector)
}

func (impl *ClientOutboundEventMaker) MakeCloseEvent(routinePoolId int64, routineId int64, uID int64, Desc string) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdClose(routinePoolId, routineId, uID, Desc)
}

func (impl *ClientOutboundEventMaker) MakeSendMessageEvent(routinePoolId int64, routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) executor.Event {
	return routineCommands.NewClientRoutineOutboundCmdMsgSend(routinePoolId, routineId, Data, OnClose, PoolKey, ChlCtx.(channel.ChannelHandlerContext), Desc)
}
