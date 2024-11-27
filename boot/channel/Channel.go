package channel

import (
	"net"
)

type Channel interface {
	IAttr

	Write(data []byte)

	Close()

	ID() string

	RemoteAddr() string
}

type IChannelError interface {
	IOReadError(err error)
	IOWriteError(err error)
}

type IChannelCallBack interface {
	Active()
	MsgReceived(data []byte)
	Inactive()
}

type ChannelBuilder interface {
	Create(conn *net.TCPConn, channelInitializer ChannelInitializer)
}
