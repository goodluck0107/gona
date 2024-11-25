package kvdb

import (
	"fmt"

	"gitee.com/andyxt/gona/kvdb/redisx"

	"github.com/go-redis/redis"
)

var optinsMap map[string]*redis.Options = make(map[string]*redis.Options, 0)

// getRedisOptions create new client options by client type
func getRedisOptions(typ string) *redis.Options {
	return optinsMap[typ]
}

func InitOptions(typ string, options *redis.Options) {
	optinsMap[typ] = options
}

func RedimoClient(typ string) *redis.Client {
	client, getErr := redisx.GetClient(getRedisOptions(typ))
	if getErr != nil {
		panic(getErr)
	}
	return client
}

// 按key获取
func GetHashTarget(typ string, k, f string) string {
	c := RedimoClient(typ)
	val, err := c.HGet(k, f).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return val
}

// 增加值
func HIncrBy(typ string, k, f string, v int64) int64 {
	c := RedimoClient(typ)
	endv, err := c.HIncrBy(k, f, v).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return endv
}

// SaveHashTarget 按key存储
func SaveHashTarget(typ string, k, f, v string) {
	c := RedimoClient(typ)
	_, err := c.HSet(k, f, v).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
}

// SaveHashMap 保存hash key的所有数据
func SaveHashMap(typ string, k string, valMap map[string]interface{}) error {
	c := RedimoClient(typ)
	_, err := c.HMSet(k, valMap).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return nil
}

// DelHashTarget 按key删除
func DelHashTarget(typ string, k, f string) {
	c := RedimoClient(typ)
	_, err := c.HDel(k, f).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
}

// GetHashAll 获取hash key的所有数据
func GetHashAll(typ string, k string) ([]string, error) {
	c := RedimoClient(typ)
	vs, err := c.HGetAll(k).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	var res []string
	for tk, tv := range vs {
		res = append(res, tk, tv)
	}
	return res, nil
}

// GetHashAllMap 获取hash key的所有数据
func GetHashAllMap(typ string, k string) (map[string]string, error) {
	c := RedimoClient(typ)
	vs, err := c.HGetAll(k).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return vs, nil
}

// 判定是否存在某个k
func Exist(typ string, k string) bool {
	c := RedimoClient(typ)
	ok, err := c.Exists(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return ok > 0
}

// 获取全局自增id
func IncrID(typ string, k string) int64 {
	c := RedimoClient(typ)
	res, err := c.Incr(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

func GetPoolAll(typ string, k string) []string {
	c := RedimoClient(typ)
	vs, err := c.HGetAll(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	var res []string
	for tk, tv := range vs {
		res = append(res, tk, tv)
	}
	return res
}

// 删除某个key
func DelKey(typ string, k string) {
	c := RedimoClient(typ)
	_, err := c.Del(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
}

// 获取字符串
func Get(typ string, k string) string {
	c := RedimoClient(typ)
	res, err := c.Get(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return fmt.Sprintf("%v", res)
}

// 设置字符串
func Set(typ string, k string, v interface{}) {
	c := RedimoClient(typ)
	_, err := c.Set(k, v, -1).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
}

// 获取set
func GetSet(typ string, k string) []string {
	c := RedimoClient(typ)
	res, err := c.SMembers(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

// 设置set
func SAdd(typ string, k, v string) int64 {
	c := RedimoClient(typ)
	res, err := c.SAdd(k, v).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

// 获取zset
func GetZSet(typ string, k string) []string {
	c := RedimoClient(typ)
	res, err := c.ZRange(k, 0, -1).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	var data []string
	for k, v := range res {
		data = append(data, fmt.Sprintf("%v", k), v)
	}
	return data
}

// 获取zset
func ZAdd(typ string, k string, f string, v float64) {
	c := RedimoClient(typ)

	_, err := c.ZAdd(k, redis.Z{
		Score:  v,
		Member: f,
	}).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
}

// All Key
func Keys(typ string, k string) []string {
	c := RedimoClient(typ)
	res, err := c.Keys(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

// Hash All Key
func HKeys(typ string, k string) []string {
	c := RedimoClient(typ)
	res, err := c.HKeys(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

// Hash 判断键值k的map中是否存在字段f
func HExist(typ string, key, field string) bool {
	c := RedimoClient(typ)
	res, err := c.HExists(key, field).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

// LAll 获取List全部数据
func LAll(typ string, k string) []string {
	c := RedimoClient(typ)
	vs, err := c.LRange(k, 0, -1).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return vs
}

// LLen 获取list长度
func LLen(typ string, k string) int64 {
	c := RedimoClient(typ)
	res, err := c.LLen(k).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return res
}

// RPush 新增list条目
func RPush(typ string, k, v string) {
	c := RedimoClient(typ)
	_, err := c.RPush(k, v).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
}

// LPop 移除list 指定数量成员
func LPop(typ string, k string, n int64) {
	c := RedimoClient(typ)
	for n > 0 {
		_, err := c.LPop(k).Result()
		if err != nil {
			panic(err)
		}
		n--
	}
}

// LREM 移除List的某个条目
func LREM(typ string, k string, v string) error {
	c := RedimoClient(typ)
	_, err := c.LRem(k, 1, v).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return nil
}
