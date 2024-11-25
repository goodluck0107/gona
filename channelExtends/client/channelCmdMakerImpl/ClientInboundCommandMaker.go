package channelCmdMakerImpl

import (
	"gitee.com/andyxt/gona/channelExtends/commandMaker/channelCmdMaker"
	"gitee.com/andyxt/gona/channelExtends/commandMaker/routineCmdMaker"
	"gitee.com/andyxt/gona/channelExtends/commands/channelCommands"
	"gitee.com/andyxt/gona/channelExtends/extends"
)

type ClientInboundCommandMaker struct {
	mEventMaker routineCmdMaker.IInboundEventMaker
}

func NewClientInboundCommandMaker(mEventMaker routineCmdMaker.IInboundEventMaker) (impl *ClientInboundCommandMaker) {
	impl = new(ClientInboundCommandMaker)
	impl.mEventMaker = mEventMaker
	return
}

// 触发异常
func (impl *ClientInboundCommandMaker) MakeExceptionCommand(ctx extends.OutterChannelHandlerContext, err error) channelCmdMaker.ICommand {
	return channelCommands.NewClientCommandException(ctx, err)
}

// 新连接
func (impl *ClientInboundCommandMaker) MakeActiveCommand(routinePoolId int64, Ctx extends.OutterChannelHandlerContext) channelCmdMaker.ICommand {
	return channelCommands.NewClientCommandActive(routinePoolId, impl.mEventMaker, Ctx)
}

// 连接中断
func (impl *ClientInboundCommandMaker) MakeInActiveCommand(routinePoolId int64, Ctx extends.OutterChannelHandlerContext) channelCmdMaker.ICommand {
	return channelCommands.NewClientCommandInActive(routinePoolId, impl.mEventMaker, Ctx)
}

// 收到消息包
func (impl *ClientInboundCommandMaker) MakeMessageReceivedCommand(routinePoolId int64, Ctx extends.OutterChannelHandlerContext, Data interface{}) channelCmdMaker.ICommand {
	return channelCommands.NewClientCommandMessageReceived(routinePoolId, impl.mEventMaker, Ctx, Data)
}
