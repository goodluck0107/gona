package listener

import (
	"gitee.com/andyxt/gona/boot/bootc/connector"
)

type IConnector interface {
	Connect(ip string, port int, params map[string]interface{})
}
type IConnAcceptor interface {
	AcceptTCP() (*TCPConnWrapper, error)
}

func Create(socketType connector.SocketType) (IConnector, IConnAcceptor) {
	acceptor := newClientTcpAcceptor()
	connector := newConnector(connector.NewTcpConnector(socketType), acceptor)
	return connector, acceptor
}
