package channel

type ChannelContext interface {
	ChannelInboundInvoker

	ChannelOutboundInvoker

	handler() (handler ChannelHandler)

	pipeline() (pipeline ChannelPipeline)

	channel() (channel Channel)

	ContextAttr() IAttr

	ID() string

	RemoteAddr() string

	/*发起写事件，消息将被送往管道处理*/
	Write(data interface{})

	/*发起关闭事件，消息将被送往管道处理*/
	Close()
}
