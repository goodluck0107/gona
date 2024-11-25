package channelInitializer

import (
	"github.com/gox-studio/gona/channelExtends/channelConsts"
	"github.com/gox-studio/gona/channelExtends/channelHandlers"
	"github.com/gox-studio/gona/channelExtends/commandMaker/channelCmdMaker"
	"github.com/gox-studio/gona/channelExtends/protocol/protocolCoder"

	"github.com/gox-studio/gona/channel"
)

type TcpChannelInitializer struct {
	routinePoolId               int64
	mInboundChannelCommandMaker channelCmdMaker.IInboundCommandMaker
	mChannelInitializerMap      map[string]channel.ChannelInitializer
	mMessageFactory             protocolCoder.IMessageFactory
}

func NewTcpChannelInitializer(routinePoolId int64,
	mInboundChannelCommandMaker channelCmdMaker.IInboundCommandMaker,
	mChannelInitializerMap map[string]channel.ChannelInitializer,
	messageFactory protocolCoder.IMessageFactory) (this *TcpChannelInitializer) {
	this = new(TcpChannelInitializer)
	this.routinePoolId = routinePoolId
	this.mInboundChannelCommandMaker = mInboundChannelCommandMaker
	this.mChannelInitializerMap = mChannelInitializerMap
	this.mMessageFactory = messageFactory
	return
}

func (initializer *TcpChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	if mChannelInitializer, ok := initializer.mChannelInitializerMap[pipeline.ContextAttr().GetString(channelConsts.ChannelTag)]; ok {
		mChannelInitializer.InitChannel(pipeline)
	}
	// UpHandler--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandler
	pipeline.AddLast("SecurityDecoder", channelHandlers.NewSecurityDecoderHandler())
	pipeline.AddLast("MessageDecoder", channelHandlers.NewMessageDecoderHandler(initializer.mMessageFactory))
	pipeline.AddLast("InBoundExecutionHandler", channelHandlers.NewInBoundExecutionHandler(initializer.routinePoolId, initializer.mInboundChannelCommandMaker))
	// DownHandler--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("SecurityEncoder", channelHandlers.NewSecurityEncoderHandler())
	pipeline.AddLast("MessageEncoder", channelHandlers.NewMessageEncoderHandler(initializer.mMessageFactory))
}
