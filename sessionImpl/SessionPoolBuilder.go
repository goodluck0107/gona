package sessionImpl

import (
	"sync"

	"gona/session"
)

type SessionPoolBuilder struct {
	mSessionPools map[int32]session.ISessionPool
	mLock         *sync.Mutex
}

func NewSessionPoolBuilder() (ret *SessionPoolBuilder) {
	ret = new(SessionPoolBuilder)
	ret.mSessionPools = make(map[int32]session.ISessionPool)
	ret.mLock = new(sync.Mutex)
	return
}

func (builder *SessionPoolBuilder) GetSessionPool(key int32) session.ISessionPool {
	defer builder.mLock.Unlock()
	builder.mLock.Lock()
	if sessionPool, ok := builder.mSessionPools[key]; ok {
		return sessionPool
	} else {
		sessionPool = session.NewSessionPool()
		builder.mSessionPools[key] = sessionPool
		return sessionPool
	}
}
