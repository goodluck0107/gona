package channel

import (
	"net"
)

// func HandleConn(conn *net.TCPConn, channelInitializer ChannelInitializer) {
// 	builder := NewSocketChannelBuilder()
// 	builder.Create(conn, channelInitializer)
// }

// func HandleConn2(conn *net.TCPConn, channelInitializer ChannelInitializer, channelParams map[string]interface{}, messageSpliter MessageSpliter) {
// 	builder := NewSocketChannelBuilder()
// 	builder.Params(channelParams)
// 	builder.MessageSpliter(messageSpliter)
// 	builder.Create(conn, channelInitializer)
// }

type SocketChannelBuilder struct {
	channelParams  map[string]interface{}
	messageSpliter MessageSpliter
}

func NewSocketChannelBuilder() (this *SocketChannelBuilder) {
	this = new(SocketChannelBuilder)
	return
}

func (builder *SocketChannelBuilder) Params(channelParams map[string]interface{}) {
	builder.channelParams = channelParams
}

func (builder *SocketChannelBuilder) MessageSpliter(messageSpliter MessageSpliter) {
	builder.messageSpliter = messageSpliter
}

func (builder *SocketChannelBuilder) Create(conn net.Conn, channelInitializer ChannelInitializer) {
	connChannel := NewSocketChannel(builder.channelParams, conn, channelInitializer)
	connChannel.SetMessageSpliter(builder.messageSpliter)
	connChannel.Start()
}
