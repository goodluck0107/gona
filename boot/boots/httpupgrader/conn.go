package httpupgrader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/goodluck0107/gona/utils"
	"io"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

const (
	channelReadLimit int32 = 256 // 256个字节
	readTimeOut      int32 = 30  // 30秒
)

// Conn is an adapter to t.Conn, which implements all t.Conn
// interface base on *websocket.Conn
type Conn struct {
	w         http.ResponseWriter
	r         *http.Request
	conn      net.Conn
	brw       *bufio.ReadWriter
	params    map[string]string
	readDone  atomic.Bool
	writeDone atomic.Bool
}

// NewConn return an initialized *WSConn
func NewConn(w http.ResponseWriter, r *http.Request, conn net.Conn, brw *bufio.ReadWriter, params map[string]string) *Conn {
	return &Conn{
		w:      w,
		r:      r,
		conn:   conn,
		brw:    brw,
		params: params,
	}
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *Conn) Read(b []byte) (int, error) {
	if c.readDone.Load() {
		for {
			_, readE := c.readUntil(16)
			if readE != nil {
				return 0, readE
			}
		}
	}
	defer c.readDone.Store(true)

	contentType := c.r.Header.Get("Content-Type")
	parts := strings.Split(contentType, ";")
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid Content-Type header: %s", contentType)
	}
	contentType = parts[0]

	switch contentType {
	case "multipart/form-data", "application/x-www-form-urlencoded":
		dataMap := make(map[string]interface{})
		query := c.r.URL.Query()
		for k, v := range query {
			dataMap[k] = v[0]
		}
		for k, v := range c.r.Form {
			dataMap[k] = v[0]
		}
		bodyData, err := json.Marshal(dataMap)
		if err != nil {
			return 0, err
		}
		n := copy(b, bodyData)
		return n, nil
	default:
		if c.r.ContentLength < 0 {
			return 0, fmt.Errorf("content length is: %d", c.r.ContentLength)
		}
		if c.r.ContentLength == 0 {
			dataMap := make(map[string]interface{})
			query := c.r.URL.Query()
			for k, v := range query {
				dataMap[k] = v[0]
			}
			bodyData, err := json.Marshal(dataMap)
			if err != nil {
				return 0, err
			}
			n := copy(b, bodyData)
			return n, nil
		}

		if c.r.ContentLength > int64(channelReadLimit) { // 防止超大内容长度导致内存溢出
			return 0, fmt.Errorf("content length too large: %d", c.r.ContentLength)
		}

		bodyData := make([]byte, int(c.r.ContentLength))
		_, err := io.ReadFull(c.brw, bodyData)
		if err != nil {
			return 0, err
		}
		n := copy(b, bodyData)
		return n, nil
	}
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *Conn) Write(b []byte) (int, error) {
	defer func() {
		c.Close()
	}()
	header := fmt.Sprintf("HTTP/1.0 200 OK\r\nContent-Length: %d\r\nX-Powered-By: Jetty\r\nAccess-Control-Max-Age: 86400\r\nAccess-Control-Allow-Credentials: true\r\nAccess-Control-Allow-Origin: *\r\nAccess-Control-Allow-Methods: GET,PUT,POST,GET,DELETE,OPTIONS\r\nAccess-Control-Allow-Headers: Origin,X-Requested-With,Content-Type,Content-Length,Accept,Authorization,X-Request-Info\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n", len(b))
	nHeader, err := c.brw.Write([]byte(header))
	if err != nil {
		return 0, err
	}
	nBody, err := c.brw.Write(b)
	if err != nil {
		return 0, err
	}
	err = c.brw.Flush()
	if err != nil {
		return 0, err
	}
	n := nHeader + nBody

	c.writeDone.Store(true)

	return n, nil
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
	return newHTTPRemoteAddr(c.conn, c.r)
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

func (c *Conn) readUntil(goal int32) (head []byte, err error) {
	var hasReadLength int32 = 0
	head = make([]byte, goal)
	for {
		var deadTime time.Time = time.Now().Add(time.Duration(readTimeOut) * time.Second)
		timeOutErr := c.conn.SetReadDeadline(deadTime)
		if timeOutErr != nil {
			err = timeOutErr
			return
		}
		i, err1 := c.conn.Read(head[hasReadLength:])
		if err1 != nil {
			err = err1
			return
		}
		if i > 0 {
			hasReadLength = hasReadLength + int32(i)
		}
		if hasReadLength >= goal {
			return
		}
	}
}

type httpRemoteAddr struct {
	conn net.Conn
	r    *http.Request
}

func newHTTPRemoteAddr(conn net.Conn, r *http.Request) *httpRemoteAddr {
	return &httpRemoteAddr{
		conn: conn,
		r:    r,
	}
}

func (hra *httpRemoteAddr) Network() string {
	return hra.conn.RemoteAddr().Network()
}

func (hra *httpRemoteAddr) String() string {
	index := strings.LastIndex(hra.r.RemoteAddr, ":")
	ip := utils.ParseIP(hra.r)
	if len(ip) > 0 {
		return fmt.Sprintf("%s%s", ip, hra.r.RemoteAddr[index:])
	}
	XForwardFor := hra.r.Header.Get("X-Forwarded-For")
	if len(XForwardFor) > 0 {
		return fmt.Sprintf("%s%s", XForwardFor, hra.r.RemoteAddr[index:])
	}
	return hra.r.RemoteAddr
}
