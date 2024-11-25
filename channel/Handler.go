package channel

type ChannelHandler interface {
	ExceptionCaught(ctx ChannelHandlerContext, err error)
}

type ChannelInboundHandler interface {
	ChannelHandler
	ChannelActive(ctx ChannelHandlerContext)(goonNext bool)
	MessageReceived(ctx ChannelHandlerContext, e interface{}) (ret interface{},goonNext bool)
	ChannelInactive(ctx ChannelHandlerContext)(goonNext bool)
}

type ChannelOutboundHandler interface {
	ChannelHandler
	Write(ctx ChannelHandlerContext, e interface{}) (ret interface{})
	Close(ctx ChannelHandlerContext)
}
