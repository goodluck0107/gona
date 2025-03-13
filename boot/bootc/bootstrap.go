package bootc

import (
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/internal/logger"
)

type bootStrap struct {
	*Options
	acceptor  listener.IConnAcceptor
	connector listener.IConnector
}

func (cbs *bootStrap) startup() error {
	go func() {
		logger.Info("开始接受服务端连接:")
		for {
			conn, err := cbs.acceptor.AcceptConn()
			if err != nil {
				logger.Error("接受服务端连接异常:", err.Error())
				continue
			}
			logger.Info("收到新的服务端连接请求")
			connParams := make(map[string]interface{})
			for k, v := range cbs.channelParams {
				connParams[k] = v
			}
			for k, v := range conn.GetParams() {
				connParams[k] = v
			}
			builder := channel.NewSocketChannelBuilder()
			builder.Params(connParams)
			builder.Create(conn.GetConn(), cbs.Initializer)
		}
	}()
	return nil
}
