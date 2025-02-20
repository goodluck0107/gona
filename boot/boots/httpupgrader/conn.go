package httpupgrader

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"gitee.com/andyxt/gona/boot"
)

// Conn is an adapter to t.Conn, which implements all t.Conn
// interface base on *websocket.Conn
type Conn struct {
	w         http.ResponseWriter
	r         *http.Request
	conn      net.Conn
	brw       *bufio.ReadWriter
	params    map[string]string
	readBuf   io.Reader
	readDone  atomic.Bool
	writeDone atomic.Bool
	startTime time.Time
}

// NewConn return an initialized *WSConn
func NewConn(w http.ResponseWriter, r *http.Request, conn net.Conn, brw *bufio.ReadWriter, params map[string]string) *Conn {
	return &Conn{
		w:         w,
		r:         r,
		conn:      conn,
		brw:       brw,
		params:    params,
		readBuf:   nil,
		startTime: time.Now(),
	}
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *Conn) Read(b []byte) (int, error) {
	if c.readDone.Load() {
		if c.writeDone.Load() {
			return 0, io.EOF
		} else if time.Now().UnixMilli()-c.startTime.UnixMilli() > 10*1000 {
			return 0, io.EOF
		}
		c.readUntil(1024 * 1024 * 1024)
		return 0, io.EOF
	}
	if c.readBuf == nil {
		var (
			err     error
			data    []byte
			dataMap = make(map[string]interface{})
		)

		query := c.r.URL.Query()
		if len(data) == 0 {
			for k, v := range query {
				dataMap[k] = v[0]
			}
		}

		contentType := strings.Split(c.r.Header.Get("Content-Type"), ";")[0]
		switch contentType {
		case "multipart/form-data", "application/x-www-form-urlencoded":
			_ = c.r.FormValue("")
			for k, v := range c.r.Form {
				dataMap[k] = v[0]
			}
			data, err = json.Marshal(dataMap)
			if err != nil {
				return 0, err
			}
		default:
			var bodyData []byte
			if c.r.ContentLength < 0 {
				return 0, fmt.Errorf("content length is: %d", c.r.ContentLength)
			}

			bodyData = make([]byte, int(c.r.ContentLength))
			_, err := io.ReadFull(c.brw, bodyData)
			if err != nil {
				return 0, err
			}

			jsonMap := make(map[string]interface{})
			if err := json.Unmarshal(bodyData, &jsonMap); err == nil {
				for k, v := range jsonMap {
					dataMap[k] = v
				}
			}
			data = bodyData
		}

		fmt.Printf("http: Type=Request,  Len=%d\n", len(data))

		buf := new(bytes.Buffer)
		_, err = buf.Write(data)
		if err != nil {
			return 0, err
		}
		c.readBuf = buf
	}

	n, err := c.readBuf.Read(b)
	if err == io.EOF {
		c.readDone.Store(true)
		return n, nil
	}

	return n, err
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *Conn) Write(b []byte) (int, error) {
	defer func() {
		c.Close()
	}()
	header := fmt.Sprintf("HTTP/1.0 200 OK\r\nContent-Length: %d\r\nX-Powered-By: Jetty\r\nAccess-Control-Max-Age: 86400\r\nAccess-Control-Allow-Credentials: true\r\nAccess-Control-Allow-Origin: *\r\nAccess-Control-Allow-Methods: GET,PUT,POST,GET,DELETE,OPTIONS\r\nAccess-Control-Allow-Headers: Origin,X-Requested-With,Content-Type,Content-Length,Accept,Authorization\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n", len(b))
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
		var deadTime time.Time = time.Now().Add(time.Duration(boot.ReadTimeOut) * time.Second)
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
	XForwardFor := hra.r.Header.Get("X-Forwarded-For")
	if len(XForwardFor) == 0 {
		return hra.r.RemoteAddr
	}
	index := strings.LastIndex(hra.r.RemoteAddr, ":")
	return fmt.Sprintf("%s%s", XForwardFor, hra.r.RemoteAddr[index:])
}
