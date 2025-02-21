package boots

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
	"github.com/gorilla/websocket"
)

type ByteOrder int8

const (
	byteOrderBigEndian    ByteOrder = 0 // 大端字节序
	byteOrderLittleEndian ByteOrder = 1 // 小端字节序
)

// Options contains some configurations for current node
type Options struct {
	TCPAddr        string
	HttpAddr       string
	TLSCertificate string // crt for tls
	TLSKey         string // key for tls
	Initializer    channel.ChannelInitializer
	MsgType        int
	Logger         logger.Logger

	ChannelParams map[string]interface{}
	ReadTimeOut   int32     // 连接读取消息超时时间
	WriteTimeOut  int32     // 连接写入消息超时时间
	ByteOrder     ByteOrder // 字节序
}

var Default = &Options{
	ChannelParams: make(map[string]interface{}),
	MsgType:       websocket.BinaryMessage,
	ByteOrder:     byteOrderBigEndian,
}

const (
	DefaultTCPAddr    string = "localhost:6829"
	DefaultClientPort        = 6829
)
