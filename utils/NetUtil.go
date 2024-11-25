package utils

import (
	"net"
)

// 设置Tcp参数
func SetConnParam(conn *net.TCPConn) {
	conn.SetNoDelay(true)
	conn.SetKeepAlive(true)
	conn.SetLinger(8)
}

// 设置Tcp参数
func SetWebSocketConnParam(conn net.Conn) {

}
