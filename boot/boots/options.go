package boots

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
	"github.com/gorilla/websocket"
)

// Options contains some configurations for current node
type Options struct {
	TCPAddr        string
	HttpAddr       string
	TLSCertificate string // crt for tls
	TLSKey         string // key for tls
	Initializer    channel.ChannelInitializer
	ChannelParams  map[string]interface{}
	MsgType        int
	Logger         logger.Logger
}

var Default = &Options{
	TCPAddr:        "",
	HttpAddr:       "",
	TLSCertificate: "",
	TLSKey:         "",
	ChannelParams:  make(map[string]interface{}),
	Initializer:    nil,
	MsgType:        DefaultMsgType,
}

const (
	DefaultTCPAddr    string = "localhost:6829"
	DefaultClientPort        = 6829
	DefaultMsgType           = websocket.BinaryMessage
)
