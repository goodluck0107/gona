package routineCmdMakerImpl

import (
	"gitee.com/andyxt/gona/channelExtends/commands/routineCommands"
	"gitee.com/andyxt/gona/channelExtends/protocol"

	"gitee.com/andyxt/gona/bootStrap/bootStrapClient/listener"
	"gitee.com/andyxt/gona/channel"
	"gitee.com/andyxt/gona/executor"
)

type ClientEventMakerImpl struct {
	connector listener.IConnector
}

func NewClientEventMakerImpl(connector listener.IConnector) (impl *ClientEventMakerImpl) {
	impl = new(ClientEventMakerImpl)
	impl.connector = connector
	return
}

// 新连接
func (impl *ClientEventMakerImpl) MakeActiveEvent(routinePoolId int64, routineId int64, Ctx channel.ChannelHandlerContext) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdActive(routinePoolId, routineId, Ctx)
}

// 连接中断
func (impl *ClientEventMakerImpl) MakeInActiveEvent(routinePoolId int64, routineId int64, Ctx channel.ChannelHandlerContext) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdInactive(routinePoolId, routineId, Ctx, impl.connector)
}

// 收到消息包
func (impl *ClientEventMakerImpl) MakeMessageReceivedEvent(routinePoolId int64, routineId int64, Data protocol.IProtocol, Ctx channel.ChannelHandlerContext) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdMsgRecv(routinePoolId, routineId, Data, Ctx)
}

// //下发消息包:OnClose是否在消息发送完毕后关闭连接
// func (impl *ClientEventMakerImpl) MakeMessageSendEvent(routinePoolId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx channel.ChannelHandlerContext, Desc string) executor.Event {
// 	return executorEvents.NewClientChannelDownMsgSendEvent(routinePoolId, Data, OnClose, PoolKey, ChlCtx, Desc)
// }

// //服务器内部消息包(跨处理流传递，保证玩家数据处理在单处理流)
// func (impl *ClientEventMakerImpl) MakeMessageStsEvent(routinePoolId int64, Data interface{}, EventQueueId int64, WaitChan chan interface{}) executor.Event {
// 	return NewChannelUpStsEvent(routinePoolId, Data, EventQueueId, WaitChan)
// }

// //发起连接
// func (impl *ClientEventMakerImpl) MakeConnectEvent(routinePoolId int64, uID int64, ip string, port int, ChannelParams map[string]interface{}) executor.Event {
// 	return NewChannelUpConnectEvent(routinePoolId, uID, ip, port, ChannelParams, impl.connector)
// }

// //关闭连接
// func (impl *ClientEventMakerImpl) MakeCloseEvent(routinePoolId int64, Id int64) executor.Event {
// 	return NewChannelUpCloseEvent(routinePoolId, Id)
// }
