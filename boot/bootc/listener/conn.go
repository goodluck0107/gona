package listener

import "net"

type ConnWrapper struct {
	conn   net.Conn
	params map[string]interface{}
}

func newConnWrapper(conn net.Conn, params map[string]interface{}) (this *ConnWrapper) {
	this = new(ConnWrapper)
	this.conn = conn
	this.params = params
	return
}

func (wrapper *ConnWrapper) GetParams() map[string]interface{} {
	params := make(map[string]interface{})
	for k, v := range wrapper.params {
		params[k] = v
	}
	return params
}

func (wrapper *ConnWrapper) GetConn() net.Conn {
	return wrapper.conn
}
