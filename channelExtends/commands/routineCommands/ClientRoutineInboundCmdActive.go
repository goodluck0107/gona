package routineCommands

import (
	"gona/channelExtends/channelConsts"
	"gona/channelExtends/extends"
	"gona/session"

	"gona/logger"
)

type ClientRoutineInboundCmdActive struct {
	routinePoolId int64
	routineId     int64
	ChlCtx        extends.OutterChannelHandlerContext
}

func NewClientRoutineInboundCmdActive(routinePoolId int64, routineId int64, ChlCtx extends.OutterChannelHandlerContext) (this *ClientRoutineInboundCmdActive) {
	this = new(ClientRoutineInboundCmdActive)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	return
}

func (activeEvent *ClientRoutineInboundCmdActive) PoolId() int64 {
	return activeEvent.routinePoolId
}

func (activeEvent *ClientRoutineInboundCmdActive) QueueId() int64 {
	return activeEvent.routineId
}

func (activeEvent *ClientRoutineInboundCmdActive) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (activeEvent *ClientRoutineInboundCmdActive) Exec() {
	logger.Debug("ClientRoutineInboundCmdActive Exec")
	uID := activeEvent.ChlCtx.ContextAttr().GetInt64(channelConsts.ChannelFireUser)
	iSession := session.GetSession(0, uID)
	if iSession == nil {
		logger.Debug("连接已经被主动关闭，新连接直接关闭")
		extends.Close(activeEvent.ChlCtx)
		return
	}
	oldChlCtx := extends.GetChlCtx(iSession)
	if oldChlCtx == nil {
		extends.ChangeChlCtx(iSession, activeEvent.ChlCtx)
		return
	}
	if !extends.ChannelContextEquals(activeEvent.ChlCtx, oldChlCtx) {
		logger.Debug("已经存在旧连接，直接关闭旧连接")
		extends.Conflict(oldChlCtx)
		extends.Close(oldChlCtx)
	}
	extends.ChangeChlCtx(iSession, activeEvent.ChlCtx)
}

type ChannelParams struct {
	Uid   int64
	Token string
}
