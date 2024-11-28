package connector

import (
	"errors"
	"net"
	"strconv"
	"time"

	"gitee.com/andyxt/gona/logger"
)

type connectEvent struct {
	ip         string
	port       int
	retryTimes int
	success    IConnectSuccess
	fail       IConnectFail
}

func newConnectEvent(ip string, port int, retryTimes int, success IConnectSuccess, fail IConnectFail) Event {
	instance := new(connectEvent)
	instance.ip = ip
	instance.port = port
	instance.retryTimes = retryTimes
	instance.success = success
	instance.fail = fail
	return instance
}

func (connectEvent *connectEvent) QueueId() int64 {
	return int64(connectEvent.port)
}

func (connectEvent *connectEvent) Exec() {
	conn, connRrr := connectEvent.createConn(connectEvent.ip, connectEvent.port)
	if connRrr != nil {
		connectEvent.fail.Handle(connRrr)
		return
	}
	connectEvent.success.Handle(conn)
}

func (e *connectEvent) createConn(ip string, port int) (net.Conn, error) {
	address := ip + ":" + strconv.Itoa(port)
	if e.retryTimes > 0 {
		for i := 0; i < e.retryTimes; i++ {
			// fmt.Println("net.Dial前 当前协程数：", runtime.NumGoroutine())
			conn, connRrr := net.Dial("tcp", address) // 底层实现时是开了新协程在处理
			// fmt.Println("net.Dial后 当前协程数：", runtime.NumGoroutine())
			if connRrr == nil {
				return conn, nil
			}
			logger.Error("tcp connect error", connRrr)
			time.Sleep(1 * time.Second)
		}
		return nil, errors.New("tcp connect fail")
	}
	for {
		// fmt.Println("net.Dial前 当前协程数：", runtime.NumGoroutine())
		conn, connRrr := net.Dial("tcp", address) // 底层实现时是开了新协程在处理
		// fmt.Println("net.Dial后 当前协程数：", runtime.NumGoroutine())
		if connRrr == nil {
			return conn, nil
		}
		logger.Error("tcp connect error", connRrr)
		time.Sleep(1 * time.Second)
	}
}
