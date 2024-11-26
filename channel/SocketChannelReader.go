package channel

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/channel/utils"
	"gitee.com/andyxt/gona/logger"
)

type SocketChannelReader struct {
	mConn             net.Conn
	mContext          Channel
	mChannelError     IChannelError
	mChannelCallBack  IChannelCallBack
	mMessageSpliter   MessageSpliter
	mChannelReadLimit int32
}

func NewSocketChannelReader(mConn net.Conn,
	mContext Channel,
	mChannelError IChannelError,
	mChannelCallBack IChannelCallBack) (this *SocketChannelReader) {
	this = new(SocketChannelReader)
	this.mConn = mConn
	this.mContext = mContext
	this.mChannelError = mChannelError
	this.mChannelCallBack = mChannelCallBack
	this.mChannelReadLimit = this.mContext.GetInt32(boot.KeyChannelReadLimit)
	if this.mChannelReadLimit <= 0 {
		this.mChannelReadLimit = boot.ChannelReadLimit
	}
	return
}
func (reader *SocketChannelReader) SetMessageSpliter(ms MessageSpliter) {
	reader.mMessageSpliter = ms
}
func (reader *SocketChannelReader) Start() {
	startChan := make(chan int, 1)
	go reader.runReadRoutine(startChan)
	<-startChan
}

func (reader *SocketChannelReader) runReadRoutine(startChan chan int) {
	defer func() {
		reader.mChannelError.IOReadError(errors.New("ReadRoutine done"))
		reader.mChannelCallBack.Inactive()

	}()
	startChan <- 1
	reader.mChannelCallBack.Active()
	for {
		if protocolData, err := reader.doRead(); err == nil {
			reader.mChannelCallBack.MsgReceived(protocolData)
		} else {
			logger.Warn("SocketChannelReader ReadRoutine", "chlCtxID=", reader.mContext.ID(), "error:", err)
			break
		}
	}
}

func (reader *SocketChannelReader) doRead() (data []byte, retErr interface{}) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	var lengthData []byte
	packageLength := reader.mMessageSpliter.GetBytesCountForMessageLength()
	if packageLength != 2 && packageLength != 4 {
		return nil, errors.New("非法包长度：" + strconv.Itoa(int(packageLength)))
	}
	lengthData, retErr = reader.readUntil(int32(packageLength))
	if retErr != nil {
		return
	}
	var messageLength int32
	if packageLength == 4 {
		messageLength = utils.ByteToInt32(lengthData)
	} else if packageLength == 2 {
		messageLength = int32(utils.ByteToInt16(lengthData))
	}
	if messageLength < 19 || messageLength >= reader.mChannelReadLimit {
		return nil, errors.New("协议非法,协议长度:" + strconv.Itoa(int(messageLength)) + ",限制长度:" + strconv.Itoa(int(reader.mChannelReadLimit)) + ",IP:" + reader.mConn.RemoteAddr().String())
	}
	var messageData []byte
	messageData, retErr = reader.readUntil(messageLength)
	if retErr != nil {
		return
	}
	data = messageData
	return
}

func (reader *SocketChannelReader) readUntil(goal int32) (head []byte, err error) {
	var hasReadLength int32 = 0
	head = make([]byte, goal)
	for {
		timeOutErr := reader.mConn.SetReadDeadline(time.Now().Add(readTimeOut))
		if timeOutErr != nil {
			err = timeOutErr
			return
		}
		i, err1 := reader.mConn.Read(head[hasReadLength:])
		if err1 != nil {
			err = err1
			return
		}
		if i > 0 {
			hasReadLength = hasReadLength + int32(i)
		}
		if hasReadLength >= goal {
			return
		}
	}
}
