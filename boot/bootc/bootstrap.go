package bootc

import (
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
)

type ClientBootStrap struct {
	*Options
	acceptor  listener.IConnAcceptor
	connector listener.IConnector
}

func (cbs *ClientBootStrap) startup() error {
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
			builder.Create(conn.GetConn(), cbs.Initializer)
		}
	}()
	return nil
}
