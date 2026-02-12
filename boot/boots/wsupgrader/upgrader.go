package wsupgrader

import (
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
)

type Upgrader struct {
	*websocket.Upgrader
}

func NewUpgrader() *Upgrader {
	return &Upgrader{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 8192,
			//WriteBufferPool: &sync.Pool{
			//	New: func() interface{} {
			//		return make([]byte, 32768)
			//	},
			// },
			HandshakeTimeout:  30 * time.Second,
			CheckOrigin:       func(_ *http.Request) bool { return true },
			EnableCompression: false,
		},
	}
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, params map[string]string, msgType int) (net.Conn, error) {
	// logger.Info(fmt.Sprintf("upgrade-websocket-connection:%+v", r.Header))
	ip := GetClientIP(r)
	conn, err := u.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return NewConn(ip, conn, params, msgType), nil
}
