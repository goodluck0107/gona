package utils

type IAttr interface {
	Get(key string) (value interface{})
	GetBool(key string) bool
	GetInt8(key string) (value int8)
	GetInt16(key string) (value int16)
	GetInt32(key string) (value int32)
	GetInt64(key string) (value int64)
	GetInt(key string) (value int)
	GetString(key string) (value string)
	Set(key string, value interface{})
	CopyToMap() map[string]interface{}
	CopyFromMap(newAttr map[string]interface{})
	Copy(newAttr IAttr)
}
