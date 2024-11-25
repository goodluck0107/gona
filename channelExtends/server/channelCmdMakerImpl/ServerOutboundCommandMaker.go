package channelCmdMakerImpl

import (
	"github.com/gox-studio/gona/channelExtends/commandMaker/channelCmdMaker"
	"github.com/gox-studio/gona/channelExtends/commands/channelCommands"

	"github.com/gox-studio/gona/channel"
)

type ServerOutboundCommandMaker struct {
}

func NewServerOutboundCommandMaker() (impl *ServerOutboundCommandMaker) {
	impl = new(ServerOutboundCommandMaker)
	return
}

// 触发异常
func (impl *ServerOutboundCommandMaker) MakeExceptionCommand(ctx channel.ChannelHandlerContext, err error) channelCmdMaker.ICommand {
	return channelCommands.NewServerCommandException(ctx, err)
}

// 请求关闭连接
func (impl *ServerOutboundCommandMaker) MakeCloseCommand(routinePoolId int32, Ctx channel.ChannelHandlerContext) channelCmdMaker.ICommand {
	return nil
}

// 下发消息包
func (impl *ServerOutboundCommandMaker) MakeMessageSendCommand(routinePoolId int32, Ctx channel.ChannelHandlerContext, Data interface{}) channelCmdMaker.ICommand {
	return nil
}
