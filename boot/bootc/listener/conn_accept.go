package listener

import "syscall"

type IConnAcceptor interface {
	AcceptConn() (*ConnWrapper, error)
}

type connAcceptor struct {
	queue chan *ConnWrapper
}

func newConnAcceptor() (this *connAcceptor) {
	this = new(connAcceptor)
	this.queue = make(chan *ConnWrapper, 10000)
	return
}

func (ca *connAcceptor) ReceiveConn(event *ConnWrapper) {
	ca.queue <- event
}

func (ca *connAcceptor) AcceptConn() (*ConnWrapper, error) {
	event, ok := <-ca.queue
	if event != nil {
		return event, nil
	}
	if !ok {
		return nil, syscall.EINVAL
	}
	return nil, syscall.EINVAL
}
