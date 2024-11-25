package routineCmdMaker

import (
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/channelExtends/protocol"

	"gitee.com/andyxt/gona/executor"
)

// Inbound
type IInboundEventMaker interface {
	//收到消息包
	MakeMessageReceivedEvent(routinePoolId int64, routineId int64, Data protocol.IProtocol, Ctx extends.OutterChannelHandlerContext) executor.Event
	//新连接
	MakeActiveEvent(routinePoolId int64, routineId int64, Ctx extends.OutterChannelHandlerContext) executor.Event
	//连接中断
	MakeInActiveEvent(routinePoolId int64, routineId int64, Ctx extends.OutterChannelHandlerContext) executor.Event
}
