package boots

import (
	"fmt"
	"net"
	"net/http"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots/httpupgrader"
	"gitee.com/andyxt/gona/boot/boots/wsupgrader"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
	"github.com/gorilla/mux"
)

type WSServerBootStrap struct {
	ip            string
	port          string // port for ws or wss
	crt           string // crt for wss
	key           string // key for wss
	channelParams map[string]interface{}
	initializer   channel.ChannelInitializer
	msgType       int
}

func NewWSServerBootStrap() (this *WSServerBootStrap) {
	this = new(WSServerBootStrap)
	this.ip = DefaultIp
	this.port = DefaultPort
	this.msgType = DefaultMsgType
	return
}

func (bootStrap *WSServerBootStrap) Params(channelParams map[string]interface{}) (ret *WSServerBootStrap) {
	bootStrap.channelParams = channelParams
	return bootStrap
}

func (bootStrap *WSServerBootStrap) MsgType(msgType int) (ret *WSServerBootStrap) {
	bootStrap.msgType = msgType
	return bootStrap
}

func (bootStrap *WSServerBootStrap) Port(port string) (ret *WSServerBootStrap) {
	bootStrap.port = port
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
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("------------new request")
		logger.Info("router", "root")
		defer func() {
			r.Body.Close()
		}()
		w.Write([]byte("OK"))
	})
	router.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bootStrap.routerHandler(params, w, r)
	})
	router.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}/{route:[A-Za-z0-9\\-]*}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bootStrap.routerHandler(params, w, r)
	})
	http.Handle("/", router)
	if bootStrap.crt != "" && bootStrap.key != "" {
		go func() {
			logger.StartUp("WSServerBootStrap 开始接受客户端wss连接:", bootStrap.port)
			err := http.ListenAndServeTLS(bootStrap.port, bootStrap.crt, bootStrap.key,
				nil)
			utils.CheckError(err)
		}()
	} else {
		go func() {
			logger.StartUp("WSServerBootStrap 开始接受客户端ws连接:", bootStrap.port)
			err := http.ListenAndServe(bootStrap.port, nil)
			utils.CheckError(err)
		}()
	}
}

func (bootStrap *WSServerBootStrap) routerHandler(params map[string]string, w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("WSServerBootStrap 收到新的客户端连接请求: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto))
	for k, v := range params {
		fmt.Println("param:", k, "=", v)
	}
	var (
		conn net.Conn
		err  error
	)
	if upgrade, ok := params["upgrade"]; ok && upgrade == "websocket" {
		logger.Info("WSServerBootStrap Upgrade websocket")
		conn, err = wsupgrader.NewUpgrader().Upgrade(w, r, params, bootStrap.msgType)
	} else {
		logger.Info("WSServerBootStrap Upgrade http")
		conn, err = httpupgrader.NewUpgrader().Upgrade(w, r, params)
	}
	if err != nil {
		logger.Error(fmt.Printf("WSServerBootStrap 接受客户端连接异常. URI=%s, Error=%s", r.RequestURI, err.Error()))
		if conn != nil {
			conn.Close()
		}
		return
	}
	connParams := make(map[string]interface{})
	for k, v := range bootStrap.channelParams {
		connParams[k] = v
	}
	connParams[boot.KeyURLPath] = r.URL.Path
	SetWebSocketConnParam(conn)
	builder := channel.NewSocketChannelBuilder()
	builder.Params(connParams)
	builder.Create(conn, bootStrap.initializer)
}

// 设置Tcp参数
func SetWebSocketConnParam(conn net.Conn) {

}

// func (bootStrap *WSServerBootStrap) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
// 	logger.Info(fmt.Sprintf("WSServerBootStrap 收到新的客户端连接请求: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, req.Proto))
// 	conn, err := wsupgrader.NewUpgrader().Upgrade(writer, req, nil, bootStrap.msgType)
// 	if err != nil {
// 		logger.Error("WSServerBootStrap 接受客户端连接异常:", err.Error())
// 		return
// 	}
// 	connParams := make(map[string]interface{})
// 	for k, v := range bootStrap.channelParams {
// 		connParams[k] = v
// 	}
// 	connParams[boot.KeyURLPath] = req.URL.Path
// 	SetWebSocketConnParam(conn)
// 	builder := channel.NewSocketChannelBuilder()
// 	builder.Params(connParams)
// 	builder.Create(conn, bootStrap.initializer)
// }
