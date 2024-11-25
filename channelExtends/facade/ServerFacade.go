package facade

import (
	"gitee.com/andyxt/gona/channelExtends/commandMaker/routineCmdMaker"
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/channelExtends/protocol"

	"gitee.com/andyxt/gona/executor"
)

type ServerFacade struct {
	routinePoolId int64
	mEventMaker   routineCmdMaker.IOutboundEventMaker
}

func NewServerFacade(routinePoolId int64, mEventMaker routineCmdMaker.IOutboundEventMaker) (facade *ServerFacade) {
	facade = new(ServerFacade)
	facade.routinePoolId = routinePoolId
	facade.mEventMaker = mEventMaker
	return
}

func (facade *ServerFacade) Connect(ip string, port int, uID int64, params map[string]interface{}) {

}

func (facade *ServerFacade) Close(uID int64, Desc string) {
}

// OnClose是否在消息发送完毕后关闭连接
func (facade *ServerFacade) SendMessage(Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) {
	executor.FireEvent(facade.mEventMaker.MakeSendMessageEvent(facade.routinePoolId, PoolKey, Data, OnClose, PoolKey, ChlCtx, Desc))
}
