package bootStrapServer

import (
	"fmt"
	"net/http"

	"gona/channel"
	"gona/channel/wsupgrader"
	"gona/logger"
	"gona/utils"
)

type WSServerBootStrap struct {
	ip             string
	port           string
	channelParams  map[string]interface{}
	initializer    channel.ChannelInitializer
	messageSpliter channel.MessageSpliter
}

func NewWSServerBootStrap() (this *WSServerBootStrap) {
	this = new(WSServerBootStrap)
	this.ip = DefaultIp
	this.port = DefaultPort
	return
}

func (bootStrap *WSServerBootStrap) Params(channelParams map[string]interface{}) (ret *WSServerBootStrap) {
	bootStrap.channelParams = channelParams
	return bootStrap
}

func (bootStrap *WSServerBootStrap) Port(port string) (ret *WSServerBootStrap) {
	bootStrap.port = port
	return bootStrap
}

func (bootStrap *WSServerBootStrap) ChannelInitializer(channelInitializer channel.ChannelInitializer) (ret *WSServerBootStrap) {
	bootStrap.initializer = channelInitializer
	return bootStrap
}

func (bootStrap *WSServerBootStrap) MessageSpliter(messageSpliter channel.MessageSpliter) (ret *WSServerBootStrap) {
	bootStrap.messageSpliter = messageSpliter
	return bootStrap
}

func (bootStrap *WSServerBootStrap) check() {
	if bootStrap.messageSpliter == nil {
		bootStrap.messageSpliter = NewDefaultMessageSpliter()
	}
	if bootStrap.channelParams == nil {
		bootStrap.channelParams = make(map[string]interface{})
	}
}

func (bootStrap *WSServerBootStrap) Listen() {
	bootStrap.check()
	go func() {
		logger.StartUp("WSServerBootStrap 开始接受客户端连接:", bootStrap.port)
		// http.Handle("/ws", websocket.Handler(bootStrap.OnWebSocket))
		err := http.ListenAndServe(bootStrap.port, bootStrap)
		utils.CheckError(err)
	}()
}

func (bootStrap *WSServerBootStrap) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	logger.Info(fmt.Sprintf("WSServerBootStrap 收到新的客户端连接请求: %s %s %s %s\n", req.RemoteAddr, req.Method, req.URL, req.Proto))
	// 把 HTTP 请求升级转换为 WebSocket 连接, 并写出 状态行 和 响应头。
	// conn 表示一个 WebSocket 连接, 调用此方法后状态行和响应头已写出, 不能再调用 writer.WriteHeader() 方法。
	conn, err := wsupgrader.NewUpgrader().Upgrade(writer, req, nil)
	if err != nil {
		logger.Error("WSServerBootStrap 接受客户端连接异常:", err.Error())
		return
	}
	utils.SetWebSocketConnParam(conn)
	builder := channel.NewSocketChannelBuilder()
	builder.Params(bootStrap.channelParams)
	builder.MessageSpliter(bootStrap.messageSpliter)
	builder.Create(conn, bootStrap.initializer)
}
