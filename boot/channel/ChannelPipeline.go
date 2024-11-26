package channel

type ChannelPipeline interface {
	ChannelInboundInvoker
	ChannelOutboundInvoker
	AddFirst(name string, handler ChannelHandler) (pipeline ChannelPipeline)
	AddLast(name string, handler ChannelHandler) (pipeline ChannelPipeline)
	channel() (channel Channel)

	ContextAttr() IAttr
}
