package boots

import (
	"fmt"
	"net"
	"net/http"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots/wsupgrader"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

type WSServerBootStrap struct {
	ip            string
	port          string // port for ws
	sport         string // port for wss
	crt           string // crt for wss
	key           string // key for wss
	channelParams map[string]interface{}
	initializer   channel.ChannelInitializer
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

func (bootStrap *WSServerBootStrap) SPort(port string) (ret *WSServerBootStrap) {
	bootStrap.sport = port
	return bootStrap
}

func (bootStrap *WSServerBootStrap) Crt(crt string) (ret *WSServerBootStrap) {
	bootStrap.crt = crt
	return bootStrap
}

func (bootStrap *WSServerBootStrap) Key(key string) (ret *WSServerBootStrap) {
	bootStrap.key = key
	return bootStrap
}

func (bootStrap *WSServerBootStrap) ChannelInitializer(channelInitializer channel.ChannelInitializer) (ret *WSServerBootStrap) {
	bootStrap.initializer = channelInitializer
	return bootStrap
}

func (bootStrap *WSServerBootStrap) check() {
	if bootStrap.channelParams == nil {
		bootStrap.channelParams = make(map[string]interface{})
	}
}

func (bootStrap *WSServerBootStrap) Listen() {
	bootStrap.check()
	if bootStrap.sport != "" && bootStrap.crt != "" && bootStrap.key != "" {
		go func() {
			logger.StartUp("WSServerBootStrap 开始接受客户端wss连接:", bootStrap.sport)
			err := http.ListenAndServeTLS(bootStrap.sport, bootStrap.crt, bootStrap.key,
				bootStrap)
			utils.CheckError(err)
		}()
	}
	go func() {
		logger.StartUp("WSServerBootStrap 开始接受客户端ws连接:", bootStrap.port)
		// http.Handle("/ws", websocket.Handler(bootStrap.OnWebSocket))
		err := http.ListenAndServe(bootStrap.port, bootStrap)
		utils.CheckError(err)
	}()
}

func (bootStrap *WSServerBootStrap) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	logger.Info(fmt.Sprintf("WSServerBootStrap 收到新的客户端连接请求: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, req.Proto))
	// 把 HTTP 请求升级转换为 WebSocket 连接, 并写出 状态行 和 响应头。
	// conn 表示一个 WebSocket 连接, 调用此方法后状态行和响应头已写出, 不能再调用 writer.WriteHeader() 方法。
	conn, err := wsupgrader.NewUpgrader().Upgrade(writer, req, nil)
	if err != nil {
		logger.Error("WSServerBootStrap 接受客户端连接异常:", err.Error())
		return
	}
	connParams := make(map[string]interface{})
	for k, v := range bootStrap.channelParams {
		connParams[k] = v
	}
	connParams[boot.KeyURLPath] = req.URL.Path
	SetWebSocketConnParam(conn)
	builder := channel.NewSocketChannelBuilder()
	builder.Params(connParams)
	builder.Create(conn, bootStrap.initializer)
}

// 设置Tcp参数
func SetWebSocketConnParam(conn net.Conn) {

}
