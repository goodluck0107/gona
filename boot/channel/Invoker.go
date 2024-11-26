package channel

type ChannelInboundInvoker interface {
	FireChannelActive()
	FireMessageReceived(event interface{})
	FireChannelInactive()
	FireExceptionCaught(err error)
}

type ChannelOutboundInvoker interface {
	FireMessageWrite(e interface{})
	FireChannelClose()
}
