package channel

import (
	"net"
	"time"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

var readTimeOut = time.Duration(30) * time.Second
var writeTimeOut = time.Duration(10) * time.Second

var packageLength int32 = 2 //消息包用几个字节表示

func Init(configFilePath string) {
	properties, parsePropertiesErr := utils.ParseProperties(configFilePath)
	if parsePropertiesErr != nil {
		logger.Error("channel.Init error:", parsePropertiesErr)
		logger.Error("channel use default properties:",
			"\nreadTimeOut=", readTimeOut,
			"\nwriteTimeOut=", writeTimeOut,
			"\npackageLength=", packageLength)
	} else {
		property, getError := properties.GetPropertyInt("readTimeOut")
		if getError != nil {
			logger.Error("channel.Init parseInt readTimeOut error:", getError)
		} else {
			readTimeOut = time.Duration(property) * time.Second
		}
		property, getError = properties.GetPropertyInt("writeTimeOut")
		if getError != nil {
			logger.Error("channel.Init parseInt writeTimeOut error:", getError)
		} else {
			writeTimeOut = time.Duration(property) * time.Second
		}
		property, getError = properties.GetPropertyInt("packageLength")
		if getError != nil {
			logger.Error("channel.Init parseInt packageLength error:", getError)
		} else {
			packageLength = int32(property)
		}
		logger.Error("channel use custom properties:",
			"\nreadTimeOut=", readTimeOut,
			"\nwriteTimeOut=", writeTimeOut,
			"\npackageLength=", packageLength)
	}
}

type Channel interface {
	IAttr

	Write(data []byte)

	Close()

	ID() string

	RemoteAddr() string
}

type IChannelError interface {
	IOReadError(err error)
	IOWriteError(err error)
}

type IChannelCallBack interface {
	Active()
	MsgReceived(data []byte)
	Inactive()
}

type ChannelBuilder interface {
	Create(conn *net.TCPConn, channelInitializer ChannelInitializer)
}
