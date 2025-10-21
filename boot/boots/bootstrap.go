package boots

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/goodluck0107/gona/boot"
	"github.com/goodluck0107/gona/boot/boots/httpupgrader"
	"github.com/goodluck0107/gona/boot/boots/wsupgrader"
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gona/internal/logger"
	"github.com/goodluck0107/gona/utils"
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
		setConnParams(conn)
		builder := channel.NewSocketChannelBuilder()
		connParams[boot.KeyConnType] = boot.ConnTypeTcp
		builder.Params(connParams)
		builder.Create(conn, bootStrap.Initializer)
	}
}

func (bootStrap *bootStrap) listenAndServeHttp() {
	router := bootStrap.configureRouter()

	addr := bootStrap.wholeInterface(bootStrap.HttpAddr)
	if bootStrap.TLSCertificate != "" && bootStrap.TLSKey != "" {
		if err := http.ListenAndServeTLS(addr, bootStrap.TLSCertificate, bootStrap.TLSKey, router); err != nil {
			utils.CheckError(err)
		}
	} else {
		if err := http.ListenAndServe(addr, router); err != nil {
			utils.CheckError(err)
		}
	}
}

func (bootStrap *bootStrap) configureRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(bootStrap.commonMiddleware)

	router.HandleFunc("/", bootStrap.rootHandler)
	router.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}", bootStrap.upgradeHandler)
	router.HandleFunc("/{upgrade:[A-Za-z0-9\\.\\-]*}/{route:.*}", bootStrap.routeHandler)

	return router
}

func (bootStrap *bootStrap) commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // 建议根据实际需求限制允许的来源
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Content-Length,Accept,Authorization,X-Request-Info")

		if r.Method == "OPTIONS" {
			if r.Body != nil {
				defer r.Body.Close()
			}
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (bootStrap *bootStrap) rootHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bootStrap.routerHandler(params, w, r)
}

func (bootStrap *bootStrap) upgradeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bootStrap.routerHandler(params, w, r)
}

func (bootStrap *bootStrap) routeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bootStrap.routerHandler(params, w, r)
}

func (bootStrap *bootStrap) routerHandler(params map[string]string, w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("收到新的http连接请求: %s %s %s %s", utils.ParseIP(r), r.Method, r.URL, r.Proto))
	if params == nil {
		logger.Error("参数为空")
		http.Error(w, "参数为空", http.StatusBadRequest)
		return
	}
	logger.Debug(fmt.Sprintf("http连接请求param: %v", params))
	connParams := applyOption(bootStrap.Options)
	initializer := bootStrap.Initializer
	msgType := bootStrap.MsgType
	hungUp := bootStrap.HttpHungup
	// 根据router重设options与initializer
	if bootStrap.RouterOptions != nil {
		var routerOptions *Options
		for _, routerOption := range bootStrap.RouterOptions {
			if routerOption.router(r.URL.Path) {
				routerOptions = routerOption.Opts
				break
			}
		}
		if routerOptions != nil {
			connParams = applyOption(routerOptions)
			initializer = routerOptions.Initializer
			msgType = routerOptions.MsgType
			hungUp = routerOptions.HttpHungup
		}
	}

	if initializer == nil {
		logger.Error("连接初始化异常:", r.URL.Path, ":", "channel initializer is nil")
		r.Body.Close()
		return
	}

	if upgrade, ok := params["upgrade"]; ok && upgrade == "websocket" {
		logger.Info("http连接请求Upgrade websocket")
		conn, err := wsupgrader.NewUpgrader().Upgrade(w, r, params, msgType)
		if err != nil {
			logger.Error(fmt.Sprintf("http连接请求Upgrade异常. uri=%s, error=%s", r.RequestURI, err.Error()))
			if c, ok := conn.(io.Closer); ok {
				c.Close()
			}
			return
		}
		connParams[boot.KeyConnType] = boot.ConnTypeWs
		connParams[boot.KeyURLPath] = r.URL.Path
		setConnParams(conn)
		builder := channel.NewSocketChannelBuilder()
		builder.Params(connParams)
		builder.Create(conn, initializer)
		return
	}
	if hungUp {
		logger.Info("http连接请求Upgrade http")
		conn, err := httpupgrader.NewUpgrader().Upgrade(w, r, params)
		if err != nil {
			logger.Error(fmt.Sprintf("http连接请求Upgrade异常. uri=%s, error=%s", r.RequestURI, err.Error()))
			if c, ok := conn.(io.Closer); ok {
				c.Close()
			}
			return
		}
		connParams[boot.KeyConnType] = boot.ConnTypeHttp
		connParams[boot.KeyURLPath] = r.URL.Path
		connParams[channel.KeyForRequest] = r
		setConnParams(conn)
		builder := channel.NewSocketChannelBuilder()
		builder.Params(connParams)
		builder.Create(conn, initializer)
		return
	}
	// handle http 请求
	connParams[boot.KeyConnType] = boot.ConnTypeHttp
	connParams[boot.KeyURLPath] = r.URL.Path
	connChannel := channel.NewHttpChannel(connParams, w, r, initializer)
	connChannel.Start()
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
	if len(opt.CustomDefine) > 0 {
		for k, v := range opt.CustomDefine {
			channelParams[k] = v
		}
	}
	return channelParams
}

func (bootStrap *bootStrap) wholeInterface(addr string) string {
	return addr[strings.Index(addr, ":"):]
}

func setConnParams(conn interface{}) {
	switch c := conn.(type) {
	case *net.TCPConn: // 设置Tcp参数
		c.SetNoDelay(true)
		c.SetKeepAlive(true)
		c.SetLinger(8)
	case net.Conn:
		// 可以添加其他类型的连接参数设置
	}
}
