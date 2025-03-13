package channel

import (
	"fmt"
	"net"
	"sync"

	"gitee.com/andyxt/gona/internal/logger"
	"gitee.com/andyxt/gona/utils"
)

const (
	ChannelKey string = "ChannelKey"
)

type SocketChannel struct {
	Attr
	mConn      net.Conn
	mPipeline  ChannelPipeline
	isInactive bool
	mReader    *SocketChannelReader
	mWriter    *SocketChannelWriter
}

func NewSocketChannel(params map[string]interface{}, conn net.Conn, channelInitializer ChannelInitializer) (this *SocketChannel) {
	return newSocketChannel(params, conn, channelInitializer)
}

func newSocketChannel(params map[string]interface{}, conn net.Conn, channelInitializer ChannelInitializer) (this *SocketChannel) {
	this = new(SocketChannel)
	this.initAttr(params)
	this.mConn = conn
	this.mPipeline = NewDefaultChannelPipeline(this)
	channelInitializer.InitChannel(this.mPipeline)
	this.mReader = NewSocketChannelReader(this.mConn, this, this, this)
	this.mWriter = NewSocketChannelWriter(this.mConn, this, this, this)
	return
}

func (chanenl *SocketChannel) initAttr(params map[string]interface{}) {
	chanenl.lock = new(sync.Mutex)
	chanenl.attr = make(map[string]interface{})
	for k, v := range params {
		chanenl.Set(k, v)
	}
	chanenl.Set(ChannelKey, utils.UUID())
}

func (chanenl *SocketChannel) ID() string {
	return chanenl.GetString(ChannelKey)
}

func (chanenl *SocketChannel) Start() {
	chanenl.startRead()
	chanenl.startWrite()
}

func (chanenl *SocketChannel) startRead() {
	chanenl.mReader.Start()
}

func (chanenl *SocketChannel) startWrite() {
	chanenl.mWriter.Start()
}

// for Channel
func (chanenl *SocketChannel) RemoteAddr() string {
	return chanenl.mConn.RemoteAddr().String()
}

// for Channel
func (chanenl *SocketChannel) Write(data []byte) {
	if !chanenl.isInactive {
		chanenl.mWriter.Write(data)
	}
}

// for Channel
func (chanenl *SocketChannel) Close() {
	if !chanenl.isInactive {
		chanenl.mWriter.Close()
	}
}

// 关闭底层网络
func (chanenl *SocketChannel) closeSocket() {
	defer func() {
		recover()
	}()
	chanenl.mConn.Close()
}

// for ChannelError
func (chanenl *SocketChannel) IOReadError(err error) {
	chlCtxKey := chanenl.ID()
	logger.Debug("SocketChannel IOReadError 1!", "chlCtxID=", chlCtxKey, " closeSocket")
	chanenl.closeSocket() //关闭底层网络
	logger.Debug("SocketChannel IOReadError 2!", "chlCtxID=", chlCtxKey, " closeWriteChan")
	chanenl.mWriter.Close() //解决写线程堵downChan
	logger.Debug("SocketChannel IOReadError 3!", "chlCtxID=", chlCtxKey, " ReadRoutine done")
}

func (chanenl *SocketChannel) IOWriteError(err error) {
	chlCtxKey := chanenl.ID()
	logger.Debug("SocketChannel IOWriteError 1!", "chlCtxID=", chlCtxKey, "closeSocket")
	chanenl.closeSocket() //关闭底层网络
	logger.Debug("SocketChannel IOWriteError 2!", "chlCtxID=", chlCtxKey, "closeWriteChan")
	chanenl.mWriter.Close() //解决写线程堵downChan
	logger.Debug("SocketChannel IOWriteError 3!", "chlCtxID=", chlCtxKey, "WriteRoutine done")
}

// for ChannelCallBack
func (chanenl *SocketChannel) Active() {
	chanenl.invokeActive()
}

func (chanenl *SocketChannel) invokeActive() {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	chanenl.mPipeline.FireChannelActive()
}

// for ChannelCallBack
func (chanenl *SocketChannel) Inactive() {
	chanenl.invokeInactive()
}

func (chanenl *SocketChannel) invokeInactive() {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	chanenl.isInactive = true
	chanenl.mPipeline.FireChannelInactive()
}

// for ChannelCallBack
func (chanenl *SocketChannel) MsgReceived(data []byte) {
	chanenl.invokeMsgReceived(data)
}

func (chanenl *SocketChannel) invokeMsgReceived(data []byte) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	chanenl.mPipeline.FireMessageReceived(data)
}
