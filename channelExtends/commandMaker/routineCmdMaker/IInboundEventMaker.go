package routineCmdMaker

import (
	"github.com/gox-studio/gona/channelExtends/extends"
	"github.com/gox-studio/gona/channelExtends/protocol"

	"github.com/gox-studio/gona/executor"
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
