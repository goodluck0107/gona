package channelCommands

import (
	"gona/channelExtends/extends"

	"gona/logger"
)

type ServerCommandActive struct {
	routinePoolId int64
	ChannelCtx    extends.OutterChannelHandlerContext
}

func NewServerCommandActive(routinePoolId int64, ChannelCtx extends.OutterChannelHandlerContext) (this *ServerCommandActive) {
	this = new(ServerCommandActive)
	this.routinePoolId = routinePoolId
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ServerCommandActive) Exec() {
	logger.Debug("ServerCommandActive Exec", extends.ChannelContextToString(event.ChannelCtx))
}
