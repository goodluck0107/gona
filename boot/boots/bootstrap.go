package boots

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots/httpupgrader"
	"gitee.com/andyxt/gona/boot/boots/wsupgrader"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
	"gitee.com/andyxt/gona/utils"
	"github.com/gorilla/mux"
)

type bootStrap struct {
	*Options
}

func (bootStrap *bootStrap) startup() error {
	if bootStrap.TCPAddr != "" {
		go bootStrap.listenAndServeTCP()
		logger.StartUp("开始接受客户端tcp连接")
	}
	if bootStrap.HttpAddr != "" {
		go bootStrap.listenAndServeHttp()
		logger.StartUp("开始接受客户端http连接")
	}
	return nil
}

func (bootStrap *bootStrap) listenAndServeTCP() {
	addr := bootStrap.wholeInterface(bootStrap.TCPAddr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		logger.Error("服务器启动异常:", addr, ":", err.Error())
		return
	}
	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logger.Error("服务器启动异常:", err.Error())
		return
	}
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logger.Error("接受客户端tcp连接异常:", err.Error())
			continue
		}
		logger.Info("收到新的客户端tcp连接请求:", conn.RemoteAddr())
		connParams := make(map[string]interface{})
		for k, v := range bootStrap.channelParams {
			connParams[k] = v
		}
		SetConnParam(conn)
		builder := channel.NewSocketChannelBuilder()
		builder.Params(connParams)
		builder.Create(conn, bootStrap.Initializer)
	}
}

func (bootStrap *bootStrap) listenAndServeHttp() {
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Powered-By", "Jetty")                                                                                    // 标识服务器端使用的技术或框架
			w.Header().Set("Content-Type", "application/json")                                                                         // 指示实际发送的数据类型的头部字段
			w.Header().Set("Access-Control-Allow-Origin", "*")                                                                         // 指定哪些网站可以参与跨源资源共享（CORS，Cross-Origin Resource Sharing）
			w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,GET,DELETE,OPTIONS")                                          // 指定允许跨域请求的 HTTP 方法
			w.Header().Set("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Content-Length,Accept,Authorization") // 指定了在跨源请求中，浏览器可以携带到服务器端的自定义请求头的列表
			if r.Method == "OPTIONS" {
				defer r.Body.Close()
				w.Header().Set("Access-Control-Allow-Credentials", "true") // 指示是否允许前端请求在跨域请求时携带认证信息（如 Cookies 和 HTTP 认证信息）
				w.Header().Set("Access-Control-Max-Age", "86400")          // 指定预检请求（preflight request）的结果（即 OPTIONS 请求的响应）可以被缓存多久
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"success:\":true}"))
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bootStrap.routerHandler(params, w, r)
	})
	router.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bootStrap.routerHandler(params, w, r)
	})
	// router.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}/{route:[A-Za-z0-9\\-]*}", func(w http.ResponseWriter, r *http.Request) {
	// 	params := mux.Vars(r)
	// 	bootStrap.routerHandler(params, w, r)
	// })
	router.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}/{route:.*}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bootStrap.routerHandler(params, w, r)
	})
	http.Handle("/", router)

	addr := bootStrap.wholeInterface(bootStrap.HttpAddr)
	if bootStrap.TLSCertificate != "" && bootStrap.TLSKey != "" {
		if err := http.ListenAndServeTLS(addr, bootStrap.TLSCertificate, bootStrap.TLSKey,
			nil); err != nil {
			utils.CheckError(err)
		}
	} else {
		if err := http.ListenAndServe(addr, nil); err != nil {
			utils.CheckError(err)
		}
	}
}

func (bootStrap *bootStrap) routerHandler(params map[string]string, w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("收到新的http连接请求: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto))
	for k, v := range params {
		logger.Info("http连接请求param:", k, "=", v)
	}
	if upgrade, ok := params["upgrade"]; ok && upgrade == "websocket" {
		logger.Info("http连接请求Upgrade websocket")
		conn, err := wsupgrader.NewUpgrader().Upgrade(w, r, params, bootStrap.MsgType)
		if err != nil {
			logger.Error(fmt.Sprintf("http连接请求Upgrade websocket异常. URI=%s, Error=%s", r.RequestURI, err.Error()))
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
		builder.Create(conn, bootStrap.Initializer)
		return
	}
	logger.Info("http连接请求Upgrade http")
	conn, err := httpupgrader.NewUpgrader().Upgrade(w, r, params)
	if err != nil {
		logger.Error(fmt.Sprintf("http连接请求Upgrade http异常. URI=%s, Error=%s", r.RequestURI, err.Error()))
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
	builder.Create(conn, bootStrap.Initializer)
}

func (bootStrap *bootStrap) wholeInterface(addr string) string {
	return addr[strings.Index(addr, ":"):]
}

// 设置Http参数
func SetWebSocketConnParam(conn net.Conn) {

}

// 设置Tcp参数
func SetConnParam(conn *net.TCPConn) {
	conn.SetNoDelay(true)
	conn.SetKeepAlive(true)
	conn.SetLinger(8)
}
