package logger

import (
	"fmt"
)

/*
uid	int	玩家uid
	atype	varchar	系统类型
	coin	int	金币
	balance	Int	当前账户金币数量
	rindex	Int	房间索引号
	siteid	Int	房间类型
	bpour	Int	房间底注
	dateline	datetime	时间
*/
func CommonSettleLog(logger *MyLog, logType string, uid int32 , coin int64, balance int64, rindex int32, siteid int8, bpour int64) {
	s := fmt.Sprintf("%d,%s,%d,%d,%d,%d,%d", uid, logType, coin, balance, rindex, siteid, bpour)
	logger.Println(s)
}

/*
uid	int	玩家uid
atype	varchar	系统类型
siteid	Int	房间类型
rindex	Int	房间索引号
bpour	Int	房间底注
seat	Int	座位号
toseat	Int	被赠送玩家座位号
touid	Int	赠送玩家uid
giftid	Int	礼物id号
coin	Int	消耗金币数量
dateline	datetime	时间
*/
func CommonGiftLog(logger *MyLog, logType string, uid int32 , siteid int8, rindex int32, seat int8, toseat  int8 , touid int32 , giftid int32, coin int64) {
	s := fmt.Sprintf("%d,%s,%d,%d,%d,%d,%d,%d,%d", uid, logType, siteid, rindex, seat, toseat, touid, giftid, coin)
	logger.Println(s)
}

/*
uid	int	玩家uid
atype	varchar	系统类型
siteid	Int	房间类型
rindex	Int	房间索引号
bpour	Int	房间底注
seat	Int	座位号
toseat	Int	被赠送玩家座位号
touid	Int	赠送玩家uid
propid	Int	礼物id号
dateline	datetime	时间
*/
func CommonToolLog(logger *MyLog, logType string, uid int32 , siteid int8, rindex int32, seat int8, toseat  int8 , touid int32 , propid int32) {
	s := fmt.Sprintf("%d,%s,%d,%d,%d,%d,%d,%d", uid, logType, siteid, rindex, seat, toseat, touid, propid)
	logger.Println(s)
}
