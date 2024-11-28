package listener

func Create() (IConnector, IConnAcceptor) {
	acceptor := newConnAcceptor()
	connector := newConnector(acceptor)
	return connector, acceptor
}
