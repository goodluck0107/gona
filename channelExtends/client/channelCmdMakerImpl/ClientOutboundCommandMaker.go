package channelCmdMakerImpl

import (
	"gona/channelExtends/commandMaker/channelCmdMaker"
	"gona/channelExtends/commands/channelCommands"

	"gona/channel"
)

type ClientOutboundCommandMaker struct {
}

func NewClientOutboundCommandMaker() (impl *ClientOutboundCommandMaker) {
	impl = new(ClientOutboundCommandMaker)
	return
}

// 触发异常
func (impl *ClientOutboundCommandMaker) MakeExceptionCommand(ctx channel.ChannelHandlerContext, err error) channelCmdMaker.ICommand {
	return channelCommands.NewClientCommandException(ctx, err)
}

// 请求关闭连接
func (impl *ClientOutboundCommandMaker) MakeCloseCommand(routinePoolId int32, Ctx channel.ChannelHandlerContext) channelCmdMaker.ICommand {
	return nil
}

// 下发消息包
func (impl *ClientOutboundCommandMaker) MakeMessageSendCommand(routinePoolId int32, Ctx channel.ChannelHandlerContext, Data interface{}) channelCmdMaker.ICommand {
	return nil
}
