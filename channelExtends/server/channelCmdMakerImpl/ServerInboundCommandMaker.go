package channelCmdMakerImpl

import (
	"gona/channelExtends/commandMaker/channelCmdMaker"
	"gona/channelExtends/commandMaker/routineCmdMaker"
	"gona/channelExtends/commands/channelCommands"

	"gona/channel"
)

type ServerInboundCommandMaker struct {
}

func NewServerInboundCommandMaker() (impl *ServerInboundCommandMaker) {
	impl = new(ServerInboundCommandMaker)
	return
}

// 触发异常
func (impl *ServerInboundCommandMaker) MakeExceptionCommand(ctx channel.ChannelHandlerContext, err error) channelCmdMaker.ICommand {
	return channelCommands.NewServerCommandException(ctx, err)
}

// 新连接
func (impl *ServerInboundCommandMaker) MakeActiveCommand(routinePoolId int64, mEventMaker routineCmdMaker.IInboundEventMaker, Ctx channel.ChannelHandlerContext) channelCmdMaker.ICommand {
	return channelCommands.NewServerCommandActive(routinePoolId, Ctx)
}

// 连接中断
func (impl *ServerInboundCommandMaker) MakeInActiveCommand(routinePoolId int64, mEventMaker routineCmdMaker.IInboundEventMaker, Ctx channel.ChannelHandlerContext) channelCmdMaker.ICommand {
	return channelCommands.NewServerCommandInActive(routinePoolId, mEventMaker, Ctx)
}

// 收到消息包
func (impl *ServerInboundCommandMaker) MakeMessageReceivedCommand(routinePoolId int64, mEventMaker routineCmdMaker.IInboundEventMaker, Ctx channel.ChannelHandlerContext, Data interface{}) channelCmdMaker.ICommand {
	return channelCommands.NewServerCommandMessageReceived(routinePoolId, mEventMaker, Ctx, Data)
}
