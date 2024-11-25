package channelCommands

import (
	"github.com/gox-studio/gona/channelExtends/commandMaker/routineCmdMaker"
	"github.com/gox-studio/gona/channelExtends/extends"
	"github.com/gox-studio/gona/channelExtends/protocol"

	"github.com/gox-studio/gona/executor"
	"github.com/gox-studio/gona/logger"
)

type ClientCommandMessageReceived struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IInboundEventMaker
	ChannelCtx    extends.OutterChannelHandlerContext
	e             interface{}
}

func NewClientCommandMessageReceived(routinePoolId int64,
	mEventMaker routineCmdMaker.IInboundEventMaker, ChannelCtx extends.OutterChannelHandlerContext, e interface{}) (this *ClientCommandMessageReceived) {
	this = new(ClientCommandMessageReceived)
	this.routinePoolId = routinePoolId
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	this.e = e
	return
}

func (event *ClientCommandMessageReceived) Exec() {
	logger.Debug("ClientCommandMessageReceived Exec", extends.ChannelContextToString(event.ChannelCtx))
	if event.ChannelCtx == nil || event.e == nil {
		return
	}
	if !extends.HasUserInfo(event.ChannelCtx) {
		logger.Debug("关闭连接：", " 关闭原因：MessageReceived but ChannelCtx is not InPool", extends.ChannelContextToString(event.ChannelCtx))
		event.ChannelCtx.Close()
		return
	}
	buf, _ := event.e.(protocol.IProtocol)
	executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(event.routinePoolId, extends.UID(event.ChannelCtx), buf, event.ChannelCtx))
}
