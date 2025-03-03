package boots_test

import (
	"fmt"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/boot/channel"
	"github.com/gorilla/websocket"
)

var WsPort int64 = 15001 // ws端口

func TestBootStrap(t *testing.T) {
	boots.Serve(
		boots.WithHttpAddr(fmt.Sprintf(":%v", WsPort),
			boots.WithRouterOption(
				"/websocket/third/pg/",
				boots.WithInitializer(NewChannelInitializer("pg")),
				boots.WithMsgType(websocket.BinaryMessage),
				boots.WithReadTimeOut(-1),
				boots.WithWriteTimeOut(-1),
				boots.WithByteOrderLittleEndian(),
				boots.WithReadLimit(512),
				boots.WithPacketBytesCount(2),
				boots.WithKeyLengthInclude(),
			),
			boots.WithRouterOption(
				"/websocket/third/jili/",
				boots.WithInitializer(NewChannelInitializer("jili")),
				boots.WithMsgType(websocket.TextMessage),
				boots.WithReadTimeOut(-1),
				boots.WithWriteTimeOut(-1),
				boots.WithByteOrderLittleEndian(),
				boots.WithReadLimit(512),
				boots.WithPacketBytesCount(2),
				boots.WithKeyLengthInclude(),
				boots.WithSkipPacketBytesCount(),
			),
		),
	)
	for {
		time.Sleep(10 * time.Second)
	}
}

type ChannelInitializer struct {
	name string
}

func (initializer *ChannelInitializer) String() string {
	return fmt.Sprintf("%vChannelInitializer", initializer.name)
}
func NewChannelInitializer(name string) (instance *ChannelInitializer) {
	instance = new(ChannelInitializer)
	instance.name = name
	return
}

func (initializer *ChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}

}
