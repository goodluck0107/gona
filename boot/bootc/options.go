package bootc

import (
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gona/internal/logger"
	"github.com/gorilla/websocket"
)

type ByteOrder int8

const (
	byteOrderBigEndian    ByteOrder = 0 // 大端字节序
	byteOrderLittleEndian ByteOrder = 1 // 小端字节序
)

// Options contains some configurations for current node
type Options struct {
	Initializer channel.ChannelInitializer
	Logger      logger.Logger

	MsgType              int
	ByteOrder            ByteOrder // 字节序
	ReadTimeOut          int32     // 连接读取消息超时时间
	WriteTimeOut         int32     // 连接写入消息超时时间
	ReadLimit            int32     // 连接读取消息长度限制
	PacketBytesCount     int32     // 消息长度占用字节数
	LengthInclude        bool      // 包长度是否包含自己的长度
	SkipPacketBytesCount bool      // 跳过包长度
	channelParams        map[string]interface{}
}

var Default = &Options{
	MsgType:              websocket.BinaryMessage,
	ByteOrder:            byteOrderBigEndian,
	ReadTimeOut:          30,    // 30秒
	WriteTimeOut:         30,    // 30秒
	ReadLimit:            256,   // 256个字节
	PacketBytesCount:     4,     // 4个字节
	LengthInclude:        false, // 包长度不包含自己的字节数
	SkipPacketBytesCount: false, // 不跳过包长度
	channelParams:        make(map[string]interface{}),
}
