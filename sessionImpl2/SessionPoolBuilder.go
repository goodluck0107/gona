package sessionImpl

import (
	"gitee.com/andyxt/gona/session"
)

type SessionPoolBuilder struct {
}

func NewSessionPoolBuilder() (ret *SessionPoolBuilder) {
	ret = new(SessionPoolBuilder)
	return
}

func (builder *SessionPoolBuilder) GetSessionPool(key int64) session.ISessionPool {
	return sessionPoolInstance
}

var sessionPoolInstance session.ISessionPool = session.NewSessionPool()

func RemoveSession(uid int64) {
	sessionPoolInstance.RemoveSession(uid)
}

func AddSession(pSession session.ISession) {
	sessionPoolInstance.AddSession(pSession)
}

func GetSession(uId int64) session.ISession {
	return sessionPoolInstance.GetSession(uId)
}

func GetSessionCount() int64 {
	return sessionPoolInstance.GetCount()
}
