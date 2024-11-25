package facade

import (
	"gona/channelExtends/extends"
	"gona/channelExtends/protocol"
)

type IServerFacade interface {
	SendMessage(Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string)
}
