package listener

import "net"

type TCPConnWrapper struct {
	conn   net.Conn
	params map[string]interface{}
}

func newTCPConnWrapper(conn net.Conn, params map[string]interface{}) (this *TCPConnWrapper) {
	this = new(TCPConnWrapper)
	this.conn = conn
	this.params = params
	return
}

func (connWrapper *TCPConnWrapper) GetParams() map[string]interface{} {
	params := make(map[string]interface{})
	for k, v := range connWrapper.params {
		params[k] = v
	}
	return params
}

func (connWrapper *TCPConnWrapper) GetConn() net.Conn {
	return connWrapper.conn
}
