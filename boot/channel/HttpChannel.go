package channel

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"gitee.com/andyxt/gona/boot/logger"
	"gitee.com/andyxt/gona/utils"
)

const (
	KeyForReqPath  = "ReqPath"  // 请求路径
	KeyForRequest  = "Request"  // *http.Request
	KeyForResponse = "Response" // http.ResponseWriter
)

type HttpChannel struct {
	Attr
	w         http.ResponseWriter
	r         *http.Request
	mPipeline ChannelPipeline
}

func NewHttpChannel(params map[string]interface{}, w http.ResponseWriter, r *http.Request, channelInitializer ChannelInitializer) (this *HttpChannel) {
	params[KeyForReqPath] = r.URL.Path // 请求路径
	params[KeyForRequest] = r          // *http.Request
	params[KeyForResponse] = w         // http.ResponseWriter
	return newHttpChannel(params, w, r, channelInitializer)
}

func newHttpChannel(params map[string]interface{}, w http.ResponseWriter, r *http.Request, channelInitializer ChannelInitializer) (this *HttpChannel) {
	this = new(HttpChannel)
	this.initAttr(params)
	this.w = w
	this.r = r
	this.mPipeline = NewDefaultChannelPipeline(this)
	channelInitializer.InitChannel(this.mPipeline)
	return
}

func (chanenl *HttpChannel) initAttr(params map[string]interface{}) {
	chanenl.lock = new(sync.Mutex)
	chanenl.attr = make(map[string]interface{})
	for k, v := range params {
		chanenl.Set(k, v)
	}
	chanenl.Set(ChannelKey, utils.UUID())
}

func (chanenl *HttpChannel) ID() string {
	return chanenl.GetString(ChannelKey)
}

func (chanenl *HttpChannel) Start() {
	defer func() {
		chanenl.r.Body.Close()
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprintf("http server error:%v", fmt.Sprint(recoverErr, string(utils.Stack(3)))))
		}
	}()
	reqBody, readErr := io.ReadAll(chanenl.r.Body)
	if readErr != nil {
		logger.Warn(fmt.Sprintf("http read error:%v", readErr))
		return
	}
	logger.Info(fmt.Sprintf("http receive reqBody:%v", reqBody))
	chanenl.mPipeline.FireMessageReceived(reqBody)
}

// for Channel
func (chanenl *HttpChannel) RemoteAddr() string {
	return chanenl.r.RemoteAddr
}

// for Channel
func (chanenl *HttpChannel) Write(data []byte) {
	chanenl.w.Write(data)
}

// for Channel
func (chanenl *HttpChannel) Close() {
	defer func() {
		recover()
	}()
	chanenl.r.Body.Close()
}

// Write 写入响应消息
// func (context *channelContext) Write(e interface{}) {
// 	defer func() {
// 		context.mReq.Body.Close()
// 		if recoverErr := recover(); recoverErr != nil {
// 			logger.Error(context.ID(), fmt.Sprint(recoverErr, string(utils.Stack(3))))
// 		}
// 	}()
// 	msg, ok := e.(message.IMessage)
// 	if !ok {
// 		logger.Error(fmt.Sprintf("%v convErr:%v", context.ID(), "write msg isnt message.IMessage"))
// 		return
// 	}
// 	respData, encodeE := msg.Encode()
// 	if encodeE != nil {
// 		logger.Error(fmt.Sprintf("%v encodeE:%v", context.ID(), encodeE))
// 		return
// 	}
// 	doWrite(context.mResp, http.StatusOK, respData)
// 	logger.Info(fmt.Sprintf("%v send respBody:%v", context.ID(), string(respData)))
// }
// // Close 发起关闭事件
// func (context *channelContext) Close() {
// 	defer func() {
// 		if recoverErr := recover(); recoverErr != nil {
// 			logger.Error(context.ID(), fmt.Sprint(recoverErr, string(utils.Stack(3))))
// 		}
// 	}()
// 	context.mReq.Body.Close()
// }

// func doWrite(Resp http.ResponseWriter, statusCode int, respData []byte) {
// 	Resp.WriteHeader(statusCode)
// 	_, writeErr := Resp.Write(respData)
// 	if writeErr != nil {
// 		logger.Error("doWrite, net write data:", respData, ", RespErr:", writeErr)
// 	}
// }
