package channel

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"gona/buffer"
	"gona/channel/utils"
	"gona/logger"
)

type SocketChannelWriter struct {
	mConn            net.Conn
	mContext         Channel
	mChannelError    IChannelError
	mChannelCallBack IChannelCallBack
	writeMsgChan     chan *WriteEvent
	mMessageSpliter  MessageSpliter
}

func NewSocketChannelWriter(mConn net.Conn,
	mContext Channel,
	mChannelError IChannelError,
	mChannelCallBack IChannelCallBack) (this *SocketChannelWriter) {
	this = new(SocketChannelWriter)
	this.mConn = mConn
	this.mContext = mContext
	this.mChannelError = mChannelError
	this.mChannelCallBack = mChannelCallBack
	this.writeMsgChan = make(chan *WriteEvent, ChannelChanSize)
	return
}

// 往写线程消息队列发送消息
func (chanenl *SocketChannelWriter) Write(data []byte) {
	defer func() {
		recover()
	}()
	chanenl.writeMsgChan <- NewWriteEvent(data, false)
}

// 关闭写线程消息队列
func (chanenl *SocketChannelWriter) Close() {
	defer func() {
		recover()
	}()
	chanenl.writeMsgChan <- NewWriteEvent(nil, true)
}
func (writer *SocketChannelWriter) SetMessageSpliter(ms MessageSpliter) {
	writer.mMessageSpliter = ms
}
func (chanenl *SocketChannelWriter) Start() {
	startChan := make(chan int, 1)
	go chanenl.runWriteRoutine(startChan)
	<-startChan
}

func (chanenl *SocketChannelWriter) runWriteRoutine(startChan chan int) {
	defer func() {
		chanenl.closeChan()
		chanenl.mChannelError.IOWriteError(errors.New("WriteRoutine done"))
	}()
	startChan <- 1
	for {
		writeEvent, ok := chanenl.getWriteData()
		if writeEvent != nil {
			data := writeEvent.data
			if data != nil {
				var messageLength = len(data)
				var lengthData []byte
				packageLength := chanenl.mMessageSpliter.GetBytesCountForMessageLength()
				if packageLength == 4 {
					lengthData = buffer.Int32ToByte(int32(messageLength))
				} else if packageLength == 2 {
					lengthData = buffer.Int16ToByte(int16(messageLength))
				} else {
					logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "error:", errors.New("非法包长度："+strconv.Itoa(int(packageLength))))
					break
				}
				if err := chanenl.doWrite(lengthData); err != nil {
					logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "error:", errors.New("非法包长度："+strconv.Itoa(int(packageLength))))
					break
				}
				if err := chanenl.doWrite(data); err != nil {
					logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "error:", errors.New("非法包长度："+strconv.Itoa(int(packageLength))))
					break
				}
			} else if writeEvent.isClose {
				logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "Received CloseMsg")
				break
			}
		} else if !ok {
			logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "DownChan is closed")
			break
		}
	}
}

func (chanenl *SocketChannelWriter) getWriteData() (data *WriteEvent, ok bool) {
	defer recover()
	data, ok = <-chanenl.writeMsgChan
	return
}

func (chanenl *SocketChannelWriter) doWrite(data []byte) (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	var goal int = len(data)
	var hasWriteLength int = 0
	for {
		timeOutErr := chanenl.mConn.SetWriteDeadline(time.Now().Add(writeTimeOut))
		if timeOutErr != nil {
			err = timeOutErr
			return
		}
		i, err1 := chanenl.mConn.Write(data)
		if err1 != nil {
			err = err1
			return
		}
		if i > 0 {
			data = data[i:]
			hasWriteLength = hasWriteLength + i
		}
		if hasWriteLength >= goal {
			return
		}
	}
}

// 关闭写线程消息队列
func (chanenl *SocketChannelWriter) closeChan() {
	defer func() {
		recover()
	}()
	close(chanenl.writeMsgChan)
}

type WriteEvent struct {
	data    []byte
	isClose bool
}

func NewWriteEvent(data []byte, isClose bool) (this *WriteEvent) {
	this = new(WriteEvent)
	this.data = data
	this.isClose = isClose
	return
}
