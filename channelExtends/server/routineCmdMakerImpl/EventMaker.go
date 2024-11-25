package routineCmdMakerImpl

import (
	"github.com/gox-studio/gona/channelExtends/commands/routineCommands"
	"github.com/gox-studio/gona/channelExtends/protocol"

	"github.com/gox-studio/gona/channel"
	"github.com/gox-studio/gona/executor"
)

type ServerEventMakerImpl struct {
}

func NewServerEventMakerImpl() (impl *ServerEventMakerImpl) {
	impl = new(ServerEventMakerImpl)
	return
}

// 新连接
func (impl *ServerEventMakerImpl) MakeActiveEvent(routinePoolId int64, routineId int64, Ctx channel.ChannelHandlerContext) executor.Event {
	return nil
}

// 连接中断
func (impl *ServerEventMakerImpl) MakeInActiveEvent(routinePoolId int64, routineId int64, Ctx channel.ChannelHandlerContext) executor.Event {
	return routineCommands.NewServerRoutineInboundCmdInactive(routinePoolId, routineId, Ctx)
}

// 收到消息包
func (impl *ServerEventMakerImpl) MakeMessageReceivedEvent(routinePoolId int64, routineId int64, Data protocol.IProtocol, Ctx channel.ChannelHandlerContext) executor.Event {
	return routineCommands.NewServerRoutineInboundCmdMsgRecv(routinePoolId, routineId, Data, Ctx)
}

// //下发消息包:OnClose是否在消息发送完毕后关闭连接
// func (impl *ServerEventMakerImpl) MakeMessageSendEvent(routinePoolId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx channel.ChannelHandlerContext, Desc string) executor.Event {
// 	return executorEvents.NewServerChannelDownMsgSendEvent(Data, OnClose, PoolKey, ChlCtx, Desc)
// }

// //服务器内部消息包(跨处理流传递，保证玩家数据处理在单处理流)
// func (impl *ServerEventMakerImpl) MakeMessageStsEvent(routinePoolId int64, Data interface{}, EventQueueId int64, WaitChan chan interface{}) executor.Event {
// 	return NewChannelUpStsEvent(Data, EventQueueId, WaitChan)
// }

// //发起连接
// func (impl *ServerEventMakerImpl) MakeConnectEvent(routinePoolId int64, uID int64, ip string, port int, ChannelParams map[string]interface{}) executor.Event {
// 	return NewChannelUpConnectEvent(ChannelParams)
// }

// //关闭连接
// func (impl *ServerEventMakerImpl) MakeCloseEvent(routinePoolId int64, Id int64) executor.Event {
// 	return NewChannelUpCloseEvent(Id)
// }
