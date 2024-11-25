package channelHandlers

import (
	"gitee.com/andyxt/gona/channelExtends/commandMaker/channelCmdMaker"
	"gitee.com/andyxt/gona/channelExtends/extends"

	"gitee.com/andyxt/gona/channel"
	"gitee.com/andyxt/gona/logger"
)

// UpBase --->
type InBoundExecutionHandler struct {
	routinePoolId        int64
	mInboundCommandMaker channelCmdMaker.IInboundCommandMaker
}

func NewInBoundExecutionHandler(routinePoolId int64, mInboundCommandMaker channelCmdMaker.IInboundCommandMaker) (this *InBoundExecutionHandler) {
	this = new(InBoundExecutionHandler)
	this.routinePoolId = routinePoolId
	this.mInboundCommandMaker = mInboundCommandMaker
	return
}

func (handler *InBoundExecutionHandler) ExceptionCaught(ctx channel.ChannelHandlerContext, err error) {
	logger.Debug("InBoundExecutionHandler ExceptionCaught", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeExceptionCommand(ctx, err).Exec()
}

func (handler *InBoundExecutionHandler) ChannelActive(ctx channel.ChannelHandlerContext) (goonNext bool) {
	logger.Debug("InBoundExecutionHandler ChannelActive", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeActiveCommand(handler.routinePoolId, ctx).Exec()
	return
}

func (handler *InBoundExecutionHandler) ChannelInactive(ctx channel.ChannelHandlerContext) (goonNext bool) {
	logger.Debug("InBoundExecutionHandler ChannelInactive", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeInActiveCommand(handler.routinePoolId, ctx).Exec()
	return
}

func (handler *InBoundExecutionHandler) MessageReceived(ctx channel.ChannelHandlerContext, e interface{}) (ret interface{}, goonNext bool) {
	logger.Debug("InBoundExecutionHandler MessageReceived", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeMessageReceivedCommand(handler.routinePoolId, ctx, e).Exec()
	return
}
