package channelCommands

import (
	"gona/channelExtends/commandMaker/routineCmdMaker"
	"gona/channelExtends/extends"

	"gona/executor"
	"gona/logger"
)

type ClientCommandInActive struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IInboundEventMaker
	ChannelCtx    extends.OutterChannelHandlerContext
}

func NewClientCommandInActive(routinePoolId int64,
	mEventMaker routineCmdMaker.IInboundEventMaker, ChannelCtx extends.OutterChannelHandlerContext) (this *ClientCommandInActive) {
	this = new(ClientCommandInActive)
	this.routinePoolId = routinePoolId
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ClientCommandInActive) Exec() {
	logger.Debug("ClientCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	if !extends.HasUserInfo(event.ChannelCtx) {
		logger.Debug("ClientCommandInActive Exec: ChannelCtx is not IsInPool", extends.ChannelContextToString(event.ChannelCtx))
		return
	}
	executor.FireEvent(event.mEventMaker.MakeInActiveEvent(event.routinePoolId, extends.UID(event.ChannelCtx), event.ChannelCtx))
}
