package channelCommands

import (
	"gona/channelExtends/extends"

	"gona/logger"
)

type ClientCommandException struct {
	ChannelCtx extends.OutterChannelHandlerContext
	err        error
}

func NewClientCommandException(ChannelCtx extends.OutterChannelHandlerContext, err error) (this *ClientCommandException) {
	this = new(ClientCommandException)
	this.ChannelCtx = ChannelCtx
	this.err = err
	return
}

func (event *ClientCommandException) Exec() {
	logger.Debug("ClientCommandException Exec", extends.ChannelContextToString(event.ChannelCtx))
	if event.ChannelCtx == nil || event.err == nil {
		return
	}
	logger.Error("关闭连接：", " 关闭原因：ClientCommandException ExceptionCaught:", event.err, extends.ChannelContextToString(event.ChannelCtx))
	extends.Close(event.ChannelCtx)
}
