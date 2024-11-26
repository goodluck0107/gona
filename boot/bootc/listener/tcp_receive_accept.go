package listener

import "syscall"

type ClientTcpAcceptor struct {
	queue chan *TCPConnWrapper
}

func newClientTcpAcceptor() (this *ClientTcpAcceptor) {
	this = new(ClientTcpAcceptor)
	this.queue = make(chan *TCPConnWrapper, 10000)
	return
}

func (clientTcpAcceptor *ClientTcpAcceptor) ReceiveTCP(event *TCPConnWrapper) {
	clientTcpAcceptor.queue <- event
}

func (clientTcpAcceptor *ClientTcpAcceptor) AcceptTCP() (*TCPConnWrapper, error) {
	event, ok := <-clientTcpAcceptor.queue
	if event != nil {
		return event, nil
	}
	if !ok {
		return nil, syscall.EINVAL
	}
	return nil, syscall.EINVAL
}
