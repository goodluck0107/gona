package boots

import (
	"fmt"
	"io"
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
		logger.StartUp("开始接受客户端tcp连接:", bootStrap.TCPAddr)
	}
	if bootStrap.HttpAddr != "" {
		go bootStrap.listenAndServeHttp()
		logger.StartUp("开始接受客户端http连接:", bootStrap.HttpAddr)
	}
	return nil
}

func (bootStrap *bootStrap) listenAndServeTCP() {
	addr := bootStrap.wholeInterface(bootStrap.TCPAddr)
	if bootStrap.Initializer == nil {
		logger.Error("服务器启动异常:", addr, ":", "channel initializer is nil")
		return
	}
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
	connParams := applyOption(bootStrap.Options)
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logger.Error("接受客户端tcp连接异常:", err.Error())
			continue
		}
		logger.Info("收到新的客户端tcp连接请求:", conn.RemoteAddr())
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
	if params == nil {
		logger.Error("参数为空")
		http.Error(w, "参数为空", http.StatusBadRequest)
		return
	}
	for k, v := range params {
		logger.Info("http连接请求param:", k, "=", v)
	}
	connParams := applyOption(bootStrap.Options)
	initializer := bootStrap.Initializer
	msgType := bootStrap.MsgType
	// 根据router重设options与initializer
	if bootStrap.RouterOptions != nil {
		var routerOptions *Options
		for routerPath, option := range bootStrap.RouterOptions {
			if strings.Contains(r.URL.Path, routerPath) {
				routerOptions = option
				break
			}
		}
		if routerOptions != nil {
			connParams = applyOption(routerOptions)
			initializer = routerOptions.Initializer
			msgType = routerOptions.MsgType
		}
	}
	var conn net.Conn
	var err error

	if upgrade, ok := params["upgrade"]; ok && upgrade == "websocket" {
		logger.Info("http连接请求Upgrade websocket")
		conn, err = wsupgrader.NewUpgrader().Upgrade(w, r, params, msgType)
	} else {
		logger.Info("http连接请求Upgrade http")
		conn, err = httpupgrader.NewUpgrader().Upgrade(w, r, params)
	}
	if err != nil {
		logger.Error(fmt.Sprintf("http连接请求Upgrade异常. uri=%s, error=%s", r.RequestURI, err.Error()))
		if c, ok := conn.(io.Closer); ok {
			c.Close()
		}
		return
	}
	if initializer == nil {
		logger.Error("连接初始化异常:", r.URL.Path, ":", "channel initializer is nil")
		if c, ok := conn.(io.Closer); ok {
			c.Close()
		}
		return
	}
	connParams[boot.KeyURLPath] = r.URL.Path
	SetWebSocketConnParam(conn)
	builder := channel.NewSocketChannelBuilder()
	builder.Params(connParams)
	builder.Create(conn, initializer)
}

func applyOption(opt *Options) map[string]any {
	channelParams := make(map[string]any)
	if opt.ByteOrder == byteOrderLittleEndian {
		channelParams[channel.KeyIsLD] = true
	}
	if opt.ReadTimeOut != 0 {
		channelParams[channel.KeyReadTimeOut] = opt.ReadTimeOut
	}
	if opt.WriteTimeOut != 0 {
		channelParams[channel.KeyWriteTimeOut] = opt.WriteTimeOut
	}
	if opt.ReadLimit > 0 {
		channelParams[channel.KeyChannelReadLimit] = opt.ReadLimit
	}
	if opt.PacketBytesCount > 0 {
		channelParams[channel.KeyPacketBytesCount] = opt.PacketBytesCount
	}
	if opt.LengthInclude {
		channelParams[channel.KeyLengthInclude] = true
	}
	if opt.SkipPacketBytesCount {
		channelParams[channel.KeySkipPacketBytesCount] = true
	}
	return channelParams
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
