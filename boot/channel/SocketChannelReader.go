package channel

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/logger"
	"gitee.com/andyxt/gona/utils"
)

type SocketChannelReader struct {
	mConn                 net.Conn
	mContext              Channel
	mChannelError         IChannelError
	mChannelCallBack      IChannelCallBack
	mChannelReadLimit     int32
	mPacketBytesCount     int32
	mReadTimeOut          int32
	mIsLD                 bool
	mLengthInclude        bool
	mSkipPacketBytesCount bool
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
	this.mPacketBytesCount = this.mContext.GetInt32(boot.KeyPacketBytesCount)
	if this.mPacketBytesCount <= 0 {
		this.mPacketBytesCount = boot.PacketBytesCount
	}
	this.mReadTimeOut = this.mContext.GetInt32(boot.KeyReadTimeOut)
	if this.mReadTimeOut == 0 {
		this.mReadTimeOut = boot.ReadTimeOut
	}
	this.mIsLD = this.mContext.GetBool(boot.KeyIsLD)
	this.mLengthInclude = this.mContext.GetBool(boot.KeyLengthInclude)
	this.mSkipPacketBytesCount = this.mContext.GetBool(boot.SkipPacketBytesCount)
	return
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
			if len(protocolData) > 0 {
				reader.mChannelCallBack.MsgReceived(protocolData)
			}
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
	if !reader.mSkipPacketBytesCount { // 不跳过消息包长度
		var messageLength int32
		var lengthData []byte
		packageLength := reader.mPacketBytesCount
		if packageLength != 2 && packageLength != 4 {
			return nil, errors.New("非法包长度：" + strconv.Itoa(int(packageLength)))
		}
		lengthData, retErr = reader.readUntil(int32(packageLength))
		if retErr != nil {
			return
		}
		logger.Debug("SocketChannelReader lengthData:", lengthData)
		if packageLength == 4 {
			if reader.mIsLD {
				messageLength = utils.ByteToInt32LD(lengthData)
			} else {
				messageLength = utils.ByteToInt32(lengthData)
			}
		} else if packageLength == 2 {
			if reader.mIsLD {
				messageLength = int32(utils.ByteToInt16LD(lengthData))
			} else {
				messageLength = int32(utils.ByteToInt16(lengthData))
			}
		}
		if messageLength < 0 || messageLength >= reader.mChannelReadLimit {
			return nil, errors.New("协议非法,协议长度:" + strconv.Itoa(int(messageLength)) + ",限制长度:" + strconv.Itoa(int(reader.mChannelReadLimit)) + ",IP:" + reader.mConn.RemoteAddr().String())
		}
		if reader.mLengthInclude {
			messageLength = messageLength - reader.mPacketBytesCount
		}
		var messageData []byte
		messageData, retErr = reader.readUntil(messageLength)
		if retErr != nil {
			return
		}
		logger.Debug("SocketChannelReader messageData:", messageData)
		data = messageData
	} else {
		var messageData []byte
		messageData, retErr = reader.readAll()
		if retErr != nil {
			return
		}
		logger.Debug("SocketChannelReader messageData:", messageData)
		data = messageData
	}
	return
}

func (reader *SocketChannelReader) readAll() (head []byte, err error) {
	ret := make([]byte, reader.mChannelReadLimit)
	var deadTime time.Time
	if reader.mReadTimeOut > 0 {
		deadTime = time.Now().Add(time.Duration(reader.mReadTimeOut) * time.Second)
	}
	timeOutErr := reader.mConn.SetReadDeadline(deadTime)
	if timeOutErr != nil {
		err = timeOutErr
		return
	}
	i, err1 := reader.mConn.Read(ret)
	if err1 != nil {
		err = err1
		return
	}
	if i > 0 {
		return ret[:i], nil
	}
	return nil, nil
}

func (reader *SocketChannelReader) readUntil(goal int32) (head []byte, err error) {
	var hasReadLength int32 = 0
	head = make([]byte, goal)
	for {
		var deadTime time.Time
		if reader.mReadTimeOut > 0 {
			deadTime = time.Now().Add(time.Duration(reader.mReadTimeOut) * time.Second)
		}
		timeOutErr := reader.mConn.SetReadDeadline(deadTime)
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
