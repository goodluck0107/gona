package utils

import "sync"

var RobertPlayerUidStart int32 = 1
var IDIncrementLock *sync.Mutex = new(sync.Mutex)

func GetAndIncrement() int32 {
	var tempUid int32 = 0
	IDIncrementLock.Lock()
	defer IDIncrementLock.Unlock()
	tempUid = RobertPlayerUidStart
	RobertPlayerUidStart=RobertPlayerUidStart+1
	return tempUid
}

/**
 * 是否是机器人用户<br>
 * 0-1999为机器人用户，所有数据库与缓存操作不需要执行<br>
 * 2000-9999为压力测试用户<br>
 * 10000-10005为测试用户
 * */
func IsRobertUser(pUid int32) bool {
	return pUid >= 50000000 && pUid < 60000000
}

/**
 * 是否是压力测试用户 0-1999为机器人用户，所有数据库与缓存操作不需要执行<br>
 * 2000-9999为压力测试用户<br>
 * 10000-10005为测试用户
 * */
func IsPressureUser(pUid int32) bool {
	return pUid >= 2000 && pUid <= 9999
}

/**
 * 是否是测试用户 0-1999为机器人用户，所有数据库与缓存操作不需要执行<br>
 * 2000-9999为压力测试用户<br>
 * 10000-10005为测试用户
 * */
func IsTestUser(pUid int32) bool {
	return pUid >= 10000 && pUid <= 10005
}
