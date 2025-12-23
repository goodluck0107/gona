package wsupgrader

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type Upgrader struct {
	*websocket.Upgrader
}

func NewUpgrader() *Upgrader {
	return &Upgrader{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:    10240,
			WriteBufferSize:   10240,
			CheckOrigin:       func(_ *http.Request) bool { return true },
			EnableCompression: false,
		},
	}
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, params map[string]string, msgType int) (net.Conn, error) {
	// logger.Info(fmt.Sprintf("upgrade-websocket-connection:%+v", r.Header))
	conn, err := u.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return NewConn(r, conn, params, msgType), nil
}
