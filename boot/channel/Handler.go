package channel

type ChannelHandler interface {
	ExceptionCaught(ctx ChannelContext, err error)
}

type ChannelInboundHandler interface {
	ChannelHandler
	ChannelActive(ctx ChannelContext) (goonNext bool)
	MessageReceived(ctx ChannelContext, e interface{}) (ret interface{}, goonNext bool)
	ChannelInactive(ctx ChannelContext) (goonNext bool)
}

type ChannelOutboundHandler interface {
	ChannelHandler
	Write(ctx ChannelContext, e interface{}) (ret interface{})
	Close(ctx ChannelContext)
}
