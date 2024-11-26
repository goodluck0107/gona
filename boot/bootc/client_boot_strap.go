package bootc

import (
	"gitee.com/andyxt/gona/boot/bootc/connector"
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
)

type ClientBootStrap struct {
	acceptor       listener.IConnAcceptor
	connector      listener.IConnector
	initializer    channel.ChannelInitializer
	messageSpliter channel.MessageSpliter
}

func NewClientBootStrap(socketType connector.SocketType) (this *ClientBootStrap) {
	this = new(ClientBootStrap)
	this.connector, this.acceptor = listener.Create(socketType)
	return this
}

func (bootStrap *ClientBootStrap) GetConnector() listener.IConnector {
	return bootStrap.connector
}

func (bootStrap *ClientBootStrap) ChannelInitializer(channelInitializer channel.ChannelInitializer) (ret *ClientBootStrap) {
	bootStrap.initializer = channelInitializer
	return bootStrap
}

func (bootStrap *ClientBootStrap) MessageSpliter(messageSpliter channel.MessageSpliter) (ret *ClientBootStrap) {
	bootStrap.messageSpliter = messageSpliter
	return bootStrap
}

func (bootStrap *ClientBootStrap) Listen() {
	go func() {
		logger.Info("开始接受服务端连接:")
		for {
			conn, err := bootStrap.acceptor.AcceptTCP()
			if err != nil {
				logger.Error("接受服务端连接异常:", err.Error())
				continue
			}
			logger.Info("收到新的服务端连接请求")
			builder := channel.NewSocketChannelBuilder()
			builder.Params(conn.GetParams())
			builder.MessageSpliter(bootStrap.messageSpliter)
			builder.Create(conn.GetConn(), bootStrap.initializer)
		}
	}()
}
