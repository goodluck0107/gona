package channelCommands

import (
	"github.com/gox-studio/gona/channelExtends/commandMaker/routineCmdMaker"
	"github.com/gox-studio/gona/channelExtends/extends"

	"github.com/gox-studio/gona/executor"
	"github.com/gox-studio/gona/logger"
)

type ServerCommandInActive struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IInboundEventMaker
	ChannelCtx    extends.OutterChannelHandlerContext
}

func NewServerCommandInActive(routinePoolId int64,
	mEventMaker routineCmdMaker.IInboundEventMaker, ChannelCtx extends.OutterChannelHandlerContext) (this *ServerCommandInActive) {
	this = new(ServerCommandInActive)
	this.routinePoolId = routinePoolId
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ServerCommandInActive) Exec() {
	logger.Debug("ServerCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	if !extends.HasUserInfo(event.ChannelCtx) {
		logger.Debug("ServerCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx), "ChannelCtx is not IsInPool")
		extends.Close(event.ChannelCtx)
		return
	}
	executor.FireEvent(event.mEventMaker.MakeInActiveEvent(event.routinePoolId, extends.UID(event.ChannelCtx), event.ChannelCtx))
}
