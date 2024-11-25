package core

const (
	result_Success        int8 = 1
	result_Failed         int8 = 2
	result_Invalid        int8 = 3 //房间信息或牌桌信息失效
	result_UnSupport      int8 = 4
	result_Error          int8 = 5
	result_TableDisbanded int8 = 6 //牌局解散

	dataTypeJson   string = "json"
	dataTypeStruct string = "struct"
)

type ExecResult struct {
	ResultType      int8
	ResultStr       string
	ResultExtendStr string
	ResultData      interface{}
	SuccessData     interface{}
	FailData        interface{}
	DataType        string
}

func (this *ExecResult) IsDataTypeJson() (bool) {
	return this.DataType == dataTypeJson
}
func (this *ExecResult) IsDataTypeStruct() (bool) {
	return this.DataType == dataTypeStruct
}
func (this *ExecResult) IsSuccess() (bool) {
	return this.ResultType == result_Success
}
func (this *ExecResult) IsFailed() (bool) {
	return this.ResultType == result_Failed
}
func (this *ExecResult) IsInvalid() (bool) {
	return this.ResultType == result_Invalid
}
func (this *ExecResult) IsUnSupport() (bool) {
	return this.ResultType == result_UnSupport
}
func (this *ExecResult) IsError() (bool) {
	return this.ResultType == result_Error
}
func (this *ExecResult) IsDisbanded() (bool) {
	return this.ResultType == result_TableDisbanded
}

func NewExecResult(resultType int8, resultStr string) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = resultType
	ret.ResultStr = resultStr
	return
}

func NewDisbandedExecResult(ResultData interface{}) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_TableDisbanded
	ret.ResultData = ResultData
	return
}

func NewFailedExecResult(resultStr string) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Failed
	ret.ResultStr = resultStr
	return
}

func NewFailedWithDataExecResultJson(resultData interface{}) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Failed
	ret.DataType = dataTypeJson
	ret.FailData = resultData
	return
}

func NewFailedWithDataExecResultStruct(resultData interface{}) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Failed
	ret.DataType = dataTypeStruct
	ret.FailData = resultData
	return
}

func NewSuccessWithDataExecResultJson(resultData interface{}) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Success
	ret.DataType = dataTypeJson
	ret.SuccessData = resultData
	return
}

func NewSuccessWithDataExecResultStruct(resultData interface{}) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Success
	ret.DataType = dataTypeStruct
	ret.SuccessData = resultData
	return
}


func NewErrorExecResult(resultStr string) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Error
	ret.ResultStr = resultStr
	return
}
func NewUnSupportExecResult(resultStr string) (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_UnSupport
	ret.ResultStr = resultStr
	return
}

func FailedExecResult() ( *ExecResult) {
	return emptyFailedExecResult
}
func SuccessExecResult() ( *ExecResult) {
	return successExecResult
}
func InvalidExecResult() ( *ExecResult) {
	return invalidExecResult
}

func (this *ExecResult) GetResultStr() string {
	return this.ResultStr
}

func newSuccessExecResult() (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Success
	return
}
func newInvalidExecResult() (ret *ExecResult) {
	ret = new(ExecResult)
	ret.ResultType = result_Invalid
	return
}

var successExecResult *ExecResult = newSuccessExecResult()
var invalidExecResult *ExecResult = newInvalidExecResult()
var emptyFailedExecResult *ExecResult = NewFailedExecResult("")
