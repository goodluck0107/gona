package boots

import (
	"github.com/gin-gonic/gin"
	"github.com/goodluck0107/gona/utils"
	"net/http"
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
	engine := gin.Default()
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

	// 2. HTTP/1.1 WebSocket 握手
	if c.Request.Method == http.MethodGet && isStdWebSocket(c.Request) {
		c.Next()
		return
	}

	// 3. HTTP/2 Extended Connect (:protocol=websocket)
	if c.Request.Method == http.MethodConnect && c.Request.Header.Get(":protocol") == "websocket" {
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
