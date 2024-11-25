package channel

type ChannelInitializer interface {
	InitChannel(pipeline ChannelPipeline)
}
