package channel

import (
	"net"
)

type SocketChannelBuilder struct {
	channelParams map[string]interface{}
}

func NewSocketChannelBuilder() (this *SocketChannelBuilder) {
	this = new(SocketChannelBuilder)
	return
}

func (builder *SocketChannelBuilder) Params(channelParams map[string]interface{}) {
	builder.channelParams = channelParams
}

func (builder *SocketChannelBuilder) Create(conn net.Conn, channelInitializer ChannelInitializer) {
	connChannel := NewSocketChannel(builder.channelParams, conn, channelInitializer)
	connChannel.Start()
}
