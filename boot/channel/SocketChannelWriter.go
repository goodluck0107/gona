package channel

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

var writeTimeOut = time.Duration(10) * time.Second

type SocketChannelWriter struct {
	mConn             net.Conn
	mContext          Channel
	mChannelError     IChannelError
	mChannelCallBack  IChannelCallBack
	writeMsgChan      chan *WriteEvent
	mPacketBytesCount int32
	mWriteTimeOut     int32
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
	this.mPacketBytesCount = this.mContext.GetInt32(boot.KeyPacketBytesCount)
	if this.mPacketBytesCount <= 0 {
		this.mPacketBytesCount = boot.PacketBytesCount
	}
	this.mWriteTimeOut = this.mContext.GetInt32(boot.KeyWriteTimeOut)
	if this.mWriteTimeOut == 0 {
		this.mWriteTimeOut = boot.WriteTimeOut
	}
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
				packageLength := chanenl.mPacketBytesCount
				if packageLength == 4 {
					lengthData = Int32ToByte(int32(messageLength))
				} else if packageLength == 2 {
					lengthData = Int16ToByte(int16(messageLength))
				} else {
					logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "error:", errors.New("非法包长度："+strconv.Itoa(int(packageLength))))
					break
				}
				logger.Debug("lengthData:", lengthData)
				if err := chanenl.doWrite(lengthData); err != nil {
					logger.Debug("SocketChannelWriter WriteRoutine", "chlCtxID=", chanenl.mContext.ID(), "error:", errors.New("非法包长度："+strconv.Itoa(int(packageLength))))
					break
				}
				logger.Debug("messageData:", data)
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
		var deadTime time.Time
		if chanenl.mWriteTimeOut > 0 {
			deadTime = time.Now().Add(time.Duration(chanenl.mWriteTimeOut) * time.Second)
		}
		timeOutErr := chanenl.mConn.SetWriteDeadline(deadTime)
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

func Int16ToByte(v int16) (buf []byte) {
	buf = make([]byte, Int16Size)
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
	return buf
}
func Int32ToByte(v int32) (buf []byte) {
	buf = make([]byte, Int32Size)
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
	return buf
}

const (
	Int16Size int32 = 2
	Int32Size int32 = 4
)
