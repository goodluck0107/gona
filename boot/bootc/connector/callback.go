package connector

import "net"

type IConnectSuccess interface {
	Handle(conn net.Conn)
}

type IConnectFail interface {
	Handle(err error)
}
