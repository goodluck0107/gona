package facade

import (
	"github.com/gox-studio/gona/channelExtends/extends"
	"github.com/gox-studio/gona/channelExtends/protocol"
)

type IServerFacade interface {
	SendMessage(Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string)
}
