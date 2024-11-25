package extends

import (
	"fmt"

	"github.com/gox-studio/gona/session"

	"github.com/gox-studio/gona/channel"
)

const sessionKeyChlCtx = "ChlCtx"

func SessionEquals(m session.ISession, n session.ISession) (ret bool) {
	return m.ID() == n.ID()
}

func SessionToString(s session.ISession) string {
	sessionID := s.ID()
	sessionUID := s.UID()
	chlCtx := GetChlCtx(s)
	chlCtxID := chlCtx.ID()
	chlCtxUID := UID(chlCtx)
	return fmt.Sprintf("sessionID=%s sessionUID=%d chlCtxID=%s chlCtxUID=%d ", sessionID, sessionUID,
		chlCtxID, chlCtxUID)
}

func ChangeChlCtx(s session.ISession, chlCtx OutterChannelHandlerContext) {
	s.SyncSet(sessionKeyChlCtx, chlCtx)
}

func GetChlCtx(s session.ISession) OutterChannelHandlerContext {
	if v := s.SyncGet(sessionKeyChlCtx); v != nil {
		contextValue, isContextValue := v.(OutterChannelHandlerContext) // Alt. non panicking version
		if isContextValue {
			return contextValue
		}
	}
	return nil
}

type OutterChannelHandlerContext interface {
	ContextAttr() channel.IAttr

	ID() string

	RemoteAddr() string

	/*发起写事件，消息将被送往管道处理*/
	Write(data interface{})

	/*发起关闭事件，消息将被送往管道处理*/
	Close()
}
