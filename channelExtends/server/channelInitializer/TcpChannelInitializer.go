package channelInitializer

import (
	"gona/channelExtends/channelHandlers"
	"gona/channelExtends/commandMaker/channelCmdMaker"
	"gona/channelExtends/protocol/protocolCoder"

	"gona/channel"
)

type TcpChannelInitializer struct {
	routinePoolId               int64
	mInboundChannelCommandMaker channelCmdMaker.IInboundCommandMaker
	mMessageFactory             protocolCoder.IMessageFactory
}

func NewTcpChannelInitializer(routinePoolId int64,
	mInboundChannelCommandMaker channelCmdMaker.IInboundCommandMaker,
	messageFactory protocolCoder.IMessageFactory) (this *TcpChannelInitializer) {
	this = new(TcpChannelInitializer)
	this.routinePoolId = routinePoolId
	this.mInboundChannelCommandMaker = mInboundChannelCommandMaker
	this.mMessageFactory = messageFactory
	return
}

func (initializer *TcpChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandler--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandler
	pipeline.AddLast("SecurityDecoder", channelHandlers.NewSecurityDecoderHandler())
	pipeline.AddLast("MessageDecoder", channelHandlers.NewMessageDecoderHandler(initializer.mMessageFactory))
	pipeline.AddLast("InBoundExecutionHandler", channelHandlers.NewInBoundExecutionHandler(initializer.routinePoolId, initializer.mInboundChannelCommandMaker))
	// DownHandler--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("SecurityEncoder", channelHandlers.NewSecurityEncoderHandler())
	pipeline.AddLast("MessageEncoder", channelHandlers.NewMessageEncoderHandler(initializer.mMessageFactory))

}
