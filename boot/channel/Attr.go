package channel

import (
	"sync"

	"gitee.com/andyxt/gona/utils/cast"
)

type Attr struct {
	lock *sync.Mutex // 扯犊子玩意，同协程都不可重入
	attr map[string]interface{}
}

func NewAttr(params map[string]interface{}) (attr *Attr) {
	attr = new(Attr)
	attr.lock = new(sync.Mutex)
	attr.attr = make(map[string]interface{})
	for k, v := range params {
		attr.Set(k, v)
	}
	return
}

func (attr *Attr) Set(key string, value interface{}) {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	attr.attr[key] = value
}

func (attr *Attr) Get(key string) (value interface{}) {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	return attr.get(key)
}

func (attr *Attr) get(key string) (value interface{}) {
	if v, ok := attr.attr[key]; ok {
		return v
	}
	return nil
}

func (attr *Attr) GetBool(key string) bool {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToBoolE(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return false
}

func (attr *Attr) GetInt8(key string) int8 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToInt8E(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return 0
}

func (attr *Attr) GetInt16(key string) int16 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToInt16E(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return 0
}

func (attr *Attr) GetInt32(key string) int32 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToInt32E(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return 0
}

func (attr *Attr) GetInt64(key string) int64 {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToInt64E(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return 0
}

func (attr *Attr) GetInt(key string) int {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToIntE(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return 0
}

func (attr *Attr) GetString(key string) string {
	defer attr.lock.Unlock()
	attr.lock.Lock()
	if v := attr.get(key); v != nil {
		castV, castE := cast.ToStringE(v) // Alt. non panicking version
		if castE == nil {
			return castV
		}
	}
	return ""
}
