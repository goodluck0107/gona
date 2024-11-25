package channelCommands

import (
	"gona/channelExtends/commandMaker/routineCmdMaker"
	"gona/channelExtends/extends"

	"gona/executor"
	"gona/il18n"
	"gona/logger"
)

type ClientCommandActive struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IInboundEventMaker
	ChannelCtx    extends.OutterChannelHandlerContext
}

func NewClientCommandActive(routinePoolId int64,
	mEventMaker routineCmdMaker.IInboundEventMaker, ChannelCtx extends.OutterChannelHandlerContext) (this *ClientCommandActive) {
	this = new(ClientCommandActive)
	this.routinePoolId = routinePoolId
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ClientCommandActive) Exec() {
	logger.Debug("ClientCommandActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	poolKey, ok := event.getFireUserID()
	if !ok {
		logger.Debug("连接激活失败：", " 失败原因：ctx.Get(ChannelId) not ok for type int64", extends.ChannelContextToString(event.ChannelCtx))
		extends.Close(event.ChannelCtx)
		return
	}
	extends.PutInUserInfo(event.ChannelCtx, poolKey, il18n.ZH_CN)
	executor.FireEvent(event.mEventMaker.MakeActiveEvent(event.routinePoolId, poolKey, event.ChannelCtx))
}

func (event *ClientCommandActive) getFireUserID() (int64, bool) {
	channelId := extends.GetFireUser(event.ChannelCtx)
	ok := channelId != -1
	return channelId, ok
}
