package utils

import (
	"log"
	"strconv"
)

//int32-10位,暂用8位,NodeId-1位,CategoryParentId-2位,CategoryId-2位,GoodId-3位
type BagIds struct {
	NodeId           int32
	CategoryParentId int32
	CategoryId       int32
	GoodId           int32
}

func NewBagIds(NodeId int32,
	CategoryParentId int32,
	CategoryId int32,
	GoodId int32) (this *BagIds) {
	this = new(BagIds)
	this.NodeId = NodeId
	this.CategoryParentId = CategoryParentId
	this.CategoryId = CategoryId
	this.GoodId = GoodId
	return
}

func NewEmptyBagIds() (this *BagIds) {
	this = new(BagIds)
	return
}

func (this *BagIds) ToId() (rlt int32) { //2,154,775,807
	rltNodeId := strconv.Itoa(int(this.NodeId))
	rltCategoryParentId := this.to2Value(this.CategoryParentId)
	rltCategoryId := this.to2Value(this.CategoryId)
	rltGoodId := this.to3Value(this.GoodId)
	rltStr := rltNodeId + rltCategoryParentId + rltCategoryId + rltGoodId
	rltInt, convertErr := strconv.Atoi(rltStr)
	if convertErr != nil {
		log.Fatalln("BagIds ToId Err:", convertErr)
		return
	}
	rlt = int32(rltInt)
	return
}

func (this *BagIds) FromId(srcId int32) { //int32-10位,暂用9位,NodeId-1位,CategoryParentId-4位,CategoryId-4位
	this.NodeId = int32(srcId / 10000000)
	this.CategoryParentId = int32(srcId % 10000000 / 100000)
	this.CategoryId = int32(srcId % 100000 / 1000)
	this.GoodId = int32(srcId % 1000)
}

func (this *BagIds) to4Value(tempId int32) (rlt string) {
	rlt = strconv.Itoa(int(tempId))
	if tempId >= 1000 {
		return rlt
	} else if tempId >= 100 {
		return "0" + rlt
	} else if tempId >= 10 {
		return "00" + rlt
	} else if tempId >= 1 {
		return "000" + rlt
	} else {
		return "0000"
	}
	return rlt
}

func (this *BagIds) to3Value(tempId int32) (rlt string) {
	rlt = strconv.Itoa(int(tempId))
	if tempId >= 100 {
		return rlt
	} else if tempId >= 10 {
		return "0" + rlt
	} else if tempId >= 1 {
		return "00" + rlt
	} else {
		return "000"
	}
	return rlt
}

func (this *BagIds) to2Value(tempId int32) (rlt string) {
	rlt = strconv.Itoa(int(tempId))
	if tempId >= 10 {
		return rlt
	} else if tempId >= 1 {
		return "0" + rlt
	} else {
		return "00"
	}
	return rlt
}
func (this *BagIds) to1Value(tempId int32) (rlt string) {
	rlt = strconv.Itoa(int(tempId))
	if tempId >= 1 {
		return rlt
	} else {
		return "0"
	}
	return rlt
}
