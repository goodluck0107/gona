package boots

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/goodluck0107/gona/utils"
	"github.com/sirupsen/logrus"
	"gitlab.yq-dev-inner.com/yq-game-developer/main-server/ck-common.git/ck-libx/utils/cast"
	"net/http"
	"strconv"
	"time"
)

func ginVars(c *gin.Context) map[string]string {
	params := make(map[string]string, len(c.Params))
	for _, p := range c.Params {
		params[p.Key] = p.Value
	}
	return params
}

func (bootStrap *bootStrap) configureRouterGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	//engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
	//	Formatter: ginLoggerJSONFormatter(),
	//	Output:    gin.DefaultWriter,
	//}))
	engine.Use(bootStrap.commonMiddlewareGin)
	engine.GET("/*path", bootStrap.rootHandlerGin)
	engine.POST("/*path", bootStrap.rootHandlerGin)
	engine.OPTIONS("/*path", bootStrap.rootHandlerGin)
	//engine.HandleFunc("/{upgrade:[A-Za-z0-9\\.]*}", bootStrap.upgradeHandlerGin)
	//engine.HandleFunc("/{upgrade:[A-Za-z0-9\\.\\-]*}/{route:.*}", bootStrap.routeHandlerGin)
	return engine
}

func (bootStrap *bootStrap) commonMiddlewareGin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Content-Length,Accept,Authorization,X-Request-Info")

	if c.Request.Method == "OPTIONS" {
		if c.Request.Body != nil {
			defer c.Request.Body.Close()
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.AbortWithStatus(http.StatusOK)
		return
	}

	// 1. POST 直接过
	if c.Request.Method == http.MethodPost {
		c.Next()
		return
	}

	if isWebSocketConn(c.Request) {
		c.Next()
		return
	}

	// 其余全部 405
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
}

func (bootStrap *bootStrap) rootHandlerGin(c *gin.Context) {
	path := c.Param("path")
	if c.Request.Method == "POST" && path == bootStrap.RouteGroup+"/ping" {
		c.String(http.StatusOK, utils.ParseIP(c.Request))
		return
	}
	params := ginVars(c)
	bootStrap.routerHandler(params, c.Writer, c.Request)
}

func (bootStrap *bootStrap) upgradeHandlerGin(c *gin.Context) {
	params := ginVars(c)
	bootStrap.routerHandler(params, c.Writer, c.Request)
}

func (bootStrap *bootStrap) routeHandlerGin(c *gin.Context) {
	params := ginVars(c)
	bootStrap.routerHandler(params, c.Writer, c.Request)
}

type outputParams struct {
	Msg        string     `json:"msg"`
	Err        string     `json:"err,omitempty"`
	Level      string     `json:"level"`
	Time       string     `json:"time"`
	HandleInfo handleInfo `json:"withFields"`
}

type handleInfo struct {
	Rid    string `json:"rid,omitempty"`
	Mid    int    `json:"mid,omitempty"`
	Uid    int64  `json:"uid,omitempty"`
	Ip     string `json:"ip"`
	Cost   string `json:"cost"`
	Mth    string `json:"method"`
	Path   string `json:"path"`
	Status int    `json:"status"`
}

func ginLoggerJSONFormatter() func(param gin.LogFormatterParams) string {
	return func(param gin.LogFormatterParams) string {
		msgId := cast.ToInt(param.Request.Header.Get("X-Request-CMD"))
		reqId := param.Request.Header.Get("X-Request-Id")
		userId := cast.ToInt64(param.Request.Header.Get("X-Inner-PlayerID"))
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		outputParam := outputParams{
			Time: param.TimeStamp.Format("2006-01-02T15:04:05.999Z07:00"),
			Msg:  "#GIN::" + strconv.Itoa(msgId),
			HandleInfo: handleInfo{
				Rid:    reqId,
				Mid:    msgId,
				Uid:    userId,
				Cost:   param.Latency.String(),
				Ip:     param.ClientIP,
				Mth:    param.Method,
				Path:   param.Path,
				Status: param.StatusCode,
			},
		}
		if param.ErrorMessage != "" {
			outputParam.Level = logrus.ErrorLevel.String()
			outputParam.Err = param.ErrorMessage
		} else if param.StatusCode != 200 {
			outputParam.Level = logrus.WarnLevel.String()
		} else {
			outputParam.Level = logrus.InfoLevel.String()
		}
		data, _ := sonic.MarshalString(&outputParam)
		return data + "\n"
	}
}

func ginLoggerTextFormatter() func(param gin.LogFormatterParams) string {
	return func(param gin.LogFormatterParams) string {
		msgId := cast.ToInt(param.Request.Header.Get("x-request-cmd"))
		reqId := cast.ToInt(param.Request.Header.Get("x-request-id"))
		userId := cast.ToInt64(param.Request.Header.Get("x-request-uid"))
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		return fmt.Sprintf("#GIN::%d %d %v |%s %3d %s| %13v | %13v | %15s |%s %-7s %s %#v\n%s",
			msgId,
			userId,
			param.TimeStamp.Format("2006-01-02T15:04:05.999Z07:00"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			reqId,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}
}
