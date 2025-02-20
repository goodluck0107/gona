package httpupgrader

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type Upgrader struct{}

func NewUpgrader() *Upgrader {
	return &Upgrader{}
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, params map[string]string) (net.Conn, error) {
	h, ok := w.(http.Hijacker)
	if !ok {
		return nil, fmt.Errorf("wrong type of response writer")
	}
	var brw *bufio.ReadWriter
	conn, brw, err := h.Hijack()
	if err != nil {
		return nil, err
	}

	c := NewConn(w, r, conn, brw, params)
	return c, nil
}
