package wsupgrader

import (
	"fmt"
	"github.com/goodluck0107/gona/utils"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Conn is an adapter to t.Conn, which implements all t.Conn
// interface base on *websocket.Conn
type Conn struct {
	ip      string
	conn    *websocket.Conn
	params  map[string]string
	typ     int // message type
	reader  io.Reader
	msgType int
}

// NewWSConn return an initialized *WSConn
func NewConn(r *http.Request, conn *websocket.Conn, params map[string]string, msgType int) *Conn {
	return &Conn{conn: conn, params: params, msgType: msgType, ip: GetClientIP(r)}
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *Conn) Read(b []byte) (int, error) {
	if c.reader == nil {
		t, r, err := c.conn.NextReader()
		if err != nil {
			return 0, err
		}
		c.typ = t
		c.reader = r
	}

	n, err := c.reader.Read(b)
	if err != nil && err != io.EOF {
		return n, err
	} else if err == io.EOF {
		_, r, err := c.conn.NextReader()
		if err != nil {
			return 0, err
		}
		c.reader = r
	}

	return n, nil
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *Conn) Write(b []byte) (int, error) {
	err := c.conn.WriteMessage(c.msgType, b) //websocket.BinaryMessage
	if err != nil {
		return 0, err
	}

	return len(b), nil
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *Conn) Close() error {
	return c.conn.Close()
}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr {
	return newWSRemoteAddr(c.conn, c.ip)
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (c *Conn) SetDeadline(t time.Time) error {
	if err := c.conn.SetReadDeadline(t); err != nil {
		return err
	}

	return c.conn.SetWriteDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

type wsRemoteAddr struct {
	conn *websocket.Conn
	ip   string
}

func newWSRemoteAddr(conn *websocket.Conn, ip string) *wsRemoteAddr {
	return &wsRemoteAddr{
		conn: conn,
		ip:   ip,
	}
}

func (ws *wsRemoteAddr) Network() string {
	return ws.conn.RemoteAddr().Network()
}

func (ws *wsRemoteAddr) String() string {
	return ws.ip
}

func GetClientIP(r *http.Request) string {
	index := strings.LastIndex(r.RemoteAddr, ":")
	ip := utils.ParseIP(r)
	if len(ip) > 0 {
		return fmt.Sprintf("%s%s", ip, r.RemoteAddr[index:])
	}
	XForwardFor := r.Header.Get("X-Forwarded-For")
	if len(XForwardFor) > 0 {
		return fmt.Sprintf("%s%s", XForwardFor, r.RemoteAddr[index:])
	}
	return r.RemoteAddr
}
