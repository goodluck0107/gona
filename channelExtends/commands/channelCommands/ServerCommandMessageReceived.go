package channelCommands

import (
	"github.com/gox-studio/gona/channelExtends/commandMaker/routineCmdMaker"
	"github.com/gox-studio/gona/channelExtends/extends"
	"github.com/gox-studio/gona/channelExtends/protocol"

	"github.com/gox-studio/gona/executor"
	"github.com/gox-studio/gona/logger"
)

type ServerCommandMessageReceived struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IInboundEventMaker
	ChannelCtx    extends.OutterChannelHandlerContext
	e             interface{}
}

func NewServerCommandMessageReceived(routinePoolId int64,
	mEventMaker routineCmdMaker.IInboundEventMaker, ChannelCtx extends.OutterChannelHandlerContext, e interface{}) (this *ServerCommandMessageReceived) {
	this = new(ServerCommandMessageReceived)
	this.routinePoolId = routinePoolId
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	this.e = e
	return
}

func (event *ServerCommandMessageReceived) Exec() {
	logger.Debug("ServerCommandMessageReceived Exec 1!", extends.ChannelContextToString(event.ChannelCtx))
	buf, _ := event.e.(protocol.IProtocol)
	if !extends.HasUserInfo(event.ChannelCtx) {
		loginData, ok := buf.(protocol.ILoginMsg)
		if !ok {
			// 尚未通过验证的连接发送任何非登陆消息都认为是非法,关闭连接
			logger.Error("ServerCommandMessageReceived Exec 2-1!", extends.ChannelContextToString(event.ChannelCtx), "关闭连接：", " 关闭原因：尚未通过验证的连接发送任何非登陆消息都认为是非法")
			extends.Close(event.ChannelCtx)
			return
		}
		uID := loginData.GetLoginUid()
		if !loginData.IsValid() {
			logger.Error("ServerCommandMessageReceived Exec 2-2!", extends.ChannelContextToString(event.ChannelCtx), "关闭连接：", " 关闭原因：发送非法登陆消息,消息内容：", loginData)
			extends.Close(event.ChannelCtx)
			return
		}
		var lngType int8 = loginData.GetLngType()
		extends.PutInUserInfo(event.ChannelCtx, uID, lngType)
		logger.Debug("ServerCommandMessageReceived Exec 2-3!", extends.ChannelContextToString(event.ChannelCtx))
	}
	logger.Debug("ServerCommandMessageReceived Exec 3-1!", extends.ChannelContextToString(event.ChannelCtx))
	eventCmd := event.mEventMaker.MakeMessageReceivedEvent(event.routinePoolId, extends.UID(event.ChannelCtx), buf, event.ChannelCtx)
	executor.FireEvent(eventCmd)
}
