package utils

import (
	"log"
	"strconv"
)

//int64-19位,暂用17位,NodeId-1位,TaskPacketId-4位,TaskDayId-4位,TaskItemId-4位,TaskRuleId-4位
type TaskIds struct {
	NodeId       int32
	TaskPacketId int32
	TaskDayId    int32
	TaskItemId   int32
	TaskRuleId   int32
}

func NewTaskIds(NodeId int32,
	TaskPacketId int32,
	TaskDayId int32,
	TaskItemId int32,
	TaskRuleId int32) (this *TaskIds) {
	this = new(TaskIds)
	this.NodeId = NodeId
	this.TaskPacketId = TaskPacketId
	this.TaskDayId = TaskDayId
	this.TaskItemId = TaskItemId
	this.TaskRuleId = TaskRuleId
	return
}

func NewEmptyTaskIds() (this *TaskIds) {
	this = new(TaskIds)
	return
}

func (this *TaskIds) ToId() (rlt int64) { //9,223,372,036,854,775,807
	rltNodeId := strconv.Itoa(int(this.NodeId))
	rltTaskPacketId := this.to4Value(this.TaskPacketId)
	rltTaskDayId := this.to4Value(this.TaskDayId)
	rltTaskItemId := this.to4Value(this.TaskItemId)
	rltTaskRuleId := this.to4Value(this.TaskRuleId)
	rltStr := rltNodeId + rltTaskPacketId + rltTaskDayId + rltTaskItemId + rltTaskRuleId
	rltInt, convertErr := strconv.Atoi(rltStr)
	if convertErr != nil {
		log.Fatalln("TaskIds ToId Err:", convertErr)
		return
	}
	rlt = int64(rltInt)
	return
}

func (this *TaskIds) FromId(srcId int64) { //int64-19位,暂用17位,NodeId-1位,TaskPacketId-4位,TaskDayId-4位,TaskItemId-4位,TaskRuleId-4位
	this.NodeId = int32(srcId / 10000000000000000)
	this.TaskPacketId = int32(srcId % 10000000000000000 / 1000000000000)
	this.TaskDayId = int32(srcId % 1000000000000 / 100000000)
	this.TaskItemId = int32(srcId % 100000000 / 10000)
	this.TaskRuleId = int32(srcId % 10000)
}

func (this *TaskIds) to4Value(tempId int32) (rlt string) {
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

func (this *TaskIds) to3Value(tempId int32) (rlt string) {
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

func (this *TaskIds) to2Value(tempId int32) (rlt string) {
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
func (this *TaskIds) to1Value(tempId int32) (rlt string) {
	rlt = strconv.Itoa(int(tempId))
	if tempId >= 1 {
		return rlt
	} else {
		return "0"
	}
	return rlt
}
