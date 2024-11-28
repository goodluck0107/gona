package bootc

import (
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
)

type ClientBootStrap struct {
	acceptor    listener.IConnAcceptor
	connector   listener.IConnector
	initializer channel.ChannelInitializer
}

func NewClientBootStrap() *ClientBootStrap {
	cbs := new(ClientBootStrap)
	cbs.connector, cbs.acceptor = listener.Create()
	return cbs
}

func (cbs *ClientBootStrap) GetConnector() listener.IConnector {
	return cbs.connector
}

func (cbs *ClientBootStrap) ChannelInitializer(channelInitializer channel.ChannelInitializer) (ret *ClientBootStrap) {
	cbs.initializer = channelInitializer
	return cbs
}

func (cbs *ClientBootStrap) Listen() {
	go func() {
		logger.Info("开始接受服务端连接:")
		for {
			conn, err := cbs.acceptor.AcceptConn()
			if err != nil {
				logger.Error("接受服务端连接异常:", err.Error())
				continue
			}
			logger.Info("收到新的服务端连接请求")
			builder := channel.NewSocketChannelBuilder()
			builder.Params(conn.GetParams())
			builder.Create(conn.GetConn(), cbs.initializer)
		}
	}()
}
