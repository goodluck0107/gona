package facade

import (
	"github.com/gox-studio/gona/channelExtends/protocol"

	"github.com/gox-studio/gona/channel"
)

type IClientFacade interface {
	Connect(ip string, port int, uID int64, params map[string]interface{})
	Close(uID int64, ChlCtx channel.ChannelHandlerContext, Desc string)
	SendMessage(Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx channel.ChannelHandlerContext, Desc string)
}
