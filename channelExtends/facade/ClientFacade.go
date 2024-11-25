package facade

import (
	"gitee.com/andyxt/gona/channelExtends/commandMaker/routineCmdMaker"
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/channelExtends/protocol"

	"gitee.com/andyxt/gona/executor"
)

type ClientFacade struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IOutboundEventMaker
}

func NewClientFacade(routinePoolId int64, mEventMaker routineCmdMaker.IOutboundEventMaker) (facade *ClientFacade) {
	facade = new(ClientFacade)
	facade.routinePoolId = routinePoolId
	facade.mEventMaker = mEventMaker
	return
}

func (facade *ClientFacade) Connect(ip string, port int, uID int64, params map[string]interface{}) {
	executor.FireEvent(facade.mEventMaker.MakeConnectEvent(facade.routinePoolId, uID, ip, port, uID, params))
}

func (facade *ClientFacade) Close(uID int64, Desc string) {
	executor.FireEvent(facade.mEventMaker.MakeCloseEvent(facade.routinePoolId, uID, uID, Desc))
}

// OnClose是否在消息发送完毕后关闭连接
func (facade *ClientFacade) SendMessage(Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) {
	executor.FireEvent(facade.mEventMaker.MakeSendMessageEvent(facade.routinePoolId, PoolKey, Data, OnClose, PoolKey, ChlCtx, Desc))
}
