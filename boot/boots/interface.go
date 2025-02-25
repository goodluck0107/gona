package boots

import (
	"fmt"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
	"github.com/mohae/deepcopy"
)

func Serve(opts ...Option) {
	opt := deepcopy.Copy(Default).(*Options)
	for _, option := range opts {
		option(opt)
	}
	if opt.Logger != nil {
		logger.Use(opt.Logger)
	}
	checkE := check(opt)
	if checkE != nil {
		panic(checkE)
	}
	bs := &bootStrap{
		Options: opt,
	}
	startupE := bs.startup()
	if startupE != nil {
		panic(startupE)
	}
}

// WithTCPAddr sets the listen address which is used to establish connection between
// cluster members. Will select an available port automatically if no member address
// setting and panic if no available port
func WithTCPAddr(addr string) Option {
	return func(opt *Options) {
		opt.TCPAddr = addr
	}
}

// WithHttpAddr sets the independent http address
func WithHttpAddr(httpAddr string) Option {
	return func(opt *Options) {
		opt.HttpAddr = httpAddr
	}
}

// WithTLSConfig sets the `key` and `certificate` of TLS
func WithTLSConfig(certificate, key string) Option {
	return func(opt *Options) {
		opt.TLSCertificate = certificate
		opt.TLSKey = key
	}
}
func WithMsgType(msgType int) Option {
	return func(opt *Options) {
		opt.MsgType = msgType
	}
}

func WithInitializer(initializer channel.ChannelInitializer) Option {
	return func(opt *Options) {
		opt.Initializer = initializer
	}
}
func WithLogger(l logger.Logger) Option {
	return func(opt *Options) {
		opt.Logger = l
	}
}

// 连接读取消息超时时间
func WithReadTimeOut(readTimeOut int32) Option {
	return func(opt *Options) {
		opt.ReadTimeOut = readTimeOut
	}
}

// 连接写入消息超时时间
func WithWriteTimeOut(writeTimeOut int32) Option {
	return func(opt *Options) {
		opt.WriteTimeOut = writeTimeOut
	}
}

// 使用小端字节序
func WithByteOrderLittleEndian() Option {
	return func(opt *Options) {
		opt.ByteOrder = byteOrderLittleEndian
	}
}

// 连接读取消息长度限制
func WithReadLimit(readLimit int32) Option {
	return func(opt *Options) {
		opt.ReadLimit = readLimit
	}
}

// 消息长度占用字节数
func WithPacketBytesCount(packetBytesCount int32) Option {
	return func(opt *Options) {
		opt.PacketBytesCount = packetBytesCount
	}
}

// 包长度是否包含自己的长度
func WithKeyLengthInclude() Option {
	return func(opt *Options) {
		opt.LengthInclude = true
	}
}

// 跳过包长度
func WithSkipPacketBytesCount() Option {
	return func(opt *Options) {
		opt.SkipPacketBytesCount = true
	}
}

type Option func(*Options)

func check(opts *Options) error {
	if opts.Initializer == nil {
		return fmt.Errorf("channel initializer is nil")
	}
	if opts.channelParams == nil {
		opts.channelParams = make(map[string]interface{})
	}
	if opts.ReadTimeOut != 0 {
		opts.channelParams[channel.KeyReadTimeOut] = opts.ReadTimeOut
	}
	if opts.WriteTimeOut != 0 {
		opts.channelParams[channel.KeyWriteTimeOut] = opts.WriteTimeOut
	}
	if opts.ByteOrder == byteOrderLittleEndian {
		opts.channelParams[channel.KeyIsLD] = true
	}
	if opts.ReadLimit > 0 {
		opts.channelParams[channel.KeyChannelReadLimit] = opts.ReadLimit
	}
	if opts.PacketBytesCount > 0 {
		opts.channelParams[channel.KeyPacketBytesCount] = opts.PacketBytesCount
	}
	if opts.LengthInclude {
		opts.channelParams[channel.KeyLengthInclude] = true
	}
	if opts.SkipPacketBytesCount {
		opts.channelParams[channel.KeySkipPacketBytesCount] = true
	}
	return nil
}
