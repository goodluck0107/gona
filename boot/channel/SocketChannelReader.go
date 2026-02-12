package channel

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/goodluck0107/gona/internal/logger"
	"github.com/goodluck0107/gona/utils"
)

const (
	KeyReadTimeOut      string = "KeyReadTimeOut"      // 连接读取消息超时时间
	KeyWriteTimeOut     string = "KeyWriteTimeOut"     // 连接写入消息超时时间
	KeyIsLD             string = "KeyIsLD"             // 是否小端
	KeyChannelReadLimit string = "KeyChannelReadLimit" // 连接读取消息长度限制

	KeyPacketBytesCount string = "KeyPacketBytesCount" // 消息长度占用字节数

	KeyLengthInclude        string = "KeyLengthInclude"        // 包长度是否包含自己的长度
	KeySkipPacketBytesCount string = "KeySkipPacketBytesCount" // 是否跳过包长度
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
	wg                    *sync.WaitGroup
}

func NewSocketChannelReader(mConn net.Conn,
	mContext Channel,
	mChannelError IChannelError,
	mChannelCallBack IChannelCallBack, wg *sync.WaitGroup) (this *SocketChannelReader) {
	this = new(SocketChannelReader)
	this.mConn = mConn
	this.mContext = mContext
	this.mChannelError = mChannelError
	this.mChannelCallBack = mChannelCallBack
	this.mChannelReadLimit = this.mContext.GetInt32(KeyChannelReadLimit)
	this.mPacketBytesCount = this.mContext.GetInt32(KeyPacketBytesCount)
	this.mReadTimeOut = this.mContext.GetInt32(KeyReadTimeOut)
	this.mIsLD = this.mContext.GetBool(KeyIsLD)
	this.mLengthInclude = this.mContext.GetBool(KeyLengthInclude)
	this.mSkipPacketBytesCount = this.mContext.GetBool(KeySkipPacketBytesCount)
	this.wg = wg
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
		reader.wg.Done()
		reader.mChannelCallBack.Inactive()
	}()
	reader.wg.Add(1)
	startChan <- 1
	reader.mChannelCallBack.Active()
	for {
		if protocolData, err := reader.doRead(); err == nil {
			if len(protocolData) > 0 {
				reader.mChannelCallBack.MsgReceived(protocolData)
			}
		} else {
			logger.Debug("SocketChannelReader ReadRoutine", "chlCtxID=", reader.mContext.ID(), "error:", err)
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
