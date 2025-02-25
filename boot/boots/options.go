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

	ChannelParams        map[string]interface{}
	ReadTimeOut          int32     // 连接读取消息超时时间
	WriteTimeOut         int32     // 连接写入消息超时时间
	ByteOrder            ByteOrder // 字节序
	ReadLimit            int32     // 连接读取消息长度限制
	PacketBytesCount     int32     // 消息长度占用字节数
	LengthInclude        bool      // 包长度是否包含自己的长度
	SkipPacketBytesCount bool      // 跳过包长度
}

var Default = &Options{
	ChannelParams:        make(map[string]interface{}),
	MsgType:              websocket.BinaryMessage,
	ByteOrder:            byteOrderBigEndian,
	ReadTimeOut:          30,    // 30秒
	WriteTimeOut:         30,    // 30秒
	ReadLimit:            256,   // 256个字节
	PacketBytesCount:     4,     // 4个字节
	LengthInclude:        false, // 包长度不包含自己的字节数
	SkipPacketBytesCount: false, // 不跳过包长度
}
