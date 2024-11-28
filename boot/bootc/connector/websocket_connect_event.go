package connector

import (
	"crypto/tls"
	"errors"
	"net"
	"net/url"
	"strconv"
	"time"

	"gitee.com/andyxt/gona/logger"
	"github.com/gorilla/websocket"
)

type websocketConnectEvent struct {
	ip         string
	port       int
	retryTimes int
	success    IConnectSuccess
	fail       IConnectFail
}

func newWebsocketConnectEvent(ip string, port int, retryTimes int, success IConnectSuccess, fail IConnectFail) Event {
	instance := new(websocketConnectEvent)
	instance.ip = ip
	instance.port = port
	instance.retryTimes = retryTimes
	instance.success = success
	instance.fail = fail
	return instance
}

func (connectEvent *websocketConnectEvent) QueueId() int64 {
	return int64(connectEvent.port)
}

func (connectEvent *websocketConnectEvent) Exec() {
	conn, connRrr := connectEvent.createConn(connectEvent.ip, connectEvent.port)
	if connRrr != nil {
		connectEvent.fail.Handle(connRrr)
		return
	}
	connectEvent.success.Handle(conn)
}

func (e *websocketConnectEvent) createConn(ip string, port int) (net.Conn, error) {
	address := ip + ":" + strconv.Itoa(port)
	if e.retryTimes > 0 {
		for i := 0; i < e.retryTimes; i++ {
			// fmt.Println("net.Dial前 当前协程数：", runtime.NumGoroutine())
			conn, connRrr := e.getWebSocketConn(address) // 底层实现时是开了新协程在处理
			// fmt.Println("net.Dial后 当前协程数：", runtime.NumGoroutine())
			if connRrr == nil {
				return conn, nil
			}
			logger.Error("ws connect error", connRrr)
			time.Sleep(1 * time.Second)
		}
		return nil, errors.New("ws connect fail")
	}
	for {
		// fmt.Println("net.Dial前 当前协程数：", runtime.NumGoroutine())
		conn, connRrr := e.getWebSocketConn(address) // 底层实现时是开了新协程在处理
		// fmt.Println("net.Dial后 当前协程数：", runtime.NumGoroutine())
		if connRrr == nil {
			return conn, nil
		}
		logger.Error("ws connect error", connRrr)
		time.Sleep(1 * time.Second)
	}
}

func (e *websocketConnectEvent) getWebSocketConn(addr string) (net.Conn, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: ""}
	dialer := websocket.DefaultDialer
	var conn *websocket.Conn
	var err error
	dialer.EnableCompression = false
	conn, _, err = dialer.Dial(u.String(), nil)
	if err != nil {
		u.Scheme = "wss"
		dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		conn, _, err = dialer.Dial(u.String(), nil)
		if err != nil {
			return nil, err
		}
	}
	return NewWSConn(conn)
}
