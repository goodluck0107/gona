package boots

import (
	"net"

	"gitee.com/andyxt/gona/channel"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

type ServerBootStrap struct {
	ip             string
	port           string
	channelParams  map[string]interface{}
	initializer    channel.ChannelInitializer
	messageSpliter channel.MessageSpliter
}

func NewServerBootStrap() (this *ServerBootStrap) {
	this = new(ServerBootStrap)
	this.ip = DefaultIp
	this.port = DefaultPort
	return
}

func (bootStrap *ServerBootStrap) Params(channelParams map[string]interface{}) (ret *ServerBootStrap) {
	bootStrap.channelParams = channelParams
	return bootStrap
}

func (bootStrap *ServerBootStrap) Port(port string) (ret *ServerBootStrap) {
	bootStrap.port = port
	return bootStrap
}

func (bootStrap *ServerBootStrap) ChannelInitializer(channelInitializer channel.ChannelInitializer) (ret *ServerBootStrap) {
	bootStrap.initializer = channelInitializer
	return bootStrap
}

func (bootStrap *ServerBootStrap) MessageSpliter(messageSpliter channel.MessageSpliter) (ret *ServerBootStrap) {
	bootStrap.messageSpliter = messageSpliter
	return bootStrap
}

func (bootStrap *ServerBootStrap) check() {
	if bootStrap.messageSpliter == nil {
		bootStrap.messageSpliter = NewDefaultMessageSpliter()
	}
	if bootStrap.channelParams == nil {
		bootStrap.channelParams = make(map[string]interface{})
	}
}

func (bootStrap *ServerBootStrap) Listen() (err error) {
	bootStrap.check()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", bootStrap.port)
	if err != nil {
		logger.Error("服务器启动异常:", bootStrap.port, ":", err.Error())
		return
	}
	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logger.Error("服务器启动异常:", err.Error())
		return
	}
	logger.StartUp("ServerBootStrap 开始接受客户端连接:", bootStrap.port)
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logger.Error("ServerBootStrap 接受客户端连接异常:", err.Error())
			continue
		}
		logger.Info("ServerBootStrap 收到新的客户端连接请求:", conn.RemoteAddr())
		utils.SetConnParam(conn)
		builder := channel.NewSocketChannelBuilder()
		builder.Params(bootStrap.channelParams)
		builder.MessageSpliter(bootStrap.messageSpliter)
		builder.Create(conn, bootStrap.initializer)
	}
}
