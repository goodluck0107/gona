package boots

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
)

func Serve(opts ...Option) {
	opt := defaultOptions()
	for _, option := range opts {
		option(opt)
	}
	if opt.Logger != nil {
		logger.Use(opt.Logger)
	}
	bs := &bootStrap{Options: opt}
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
func WithHttpAddr(httpAddr string, opts ...*RouterOption) Option {
	return func(opt *Options) {
		opt.HttpAddr = httpAddr
		opt.RouterOptions = make([]*RouterOption, 0)
		for _, v := range opts {
			opt.RouterOptions = append(opt.RouterOptions, v)
		}
	}
}

func WithRouterOption(router Router, opts ...Option) *RouterOption {
	opt := defaultOptions()
	for _, v := range opts {
		v(opt)
	}
	return &RouterOption{router: router, Opts: opt}
}

// WithTLSConfig sets the `key` and `certificate` of TLS
func WithTLSConfig(certificate, key string) Option {
	return func(opt *Options) {
		opt.TLSCertificate = certificate
		opt.TLSKey = key
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

func WithMsgType(msgType int) Option {
	return func(opt *Options) {
		opt.MsgType = msgType
	}
}

func WithHttpHungup() Option {
	return func(opt *Options) {
		opt.HttpHungup = true
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

func WithCustom(key string, value any) Option {
	return func(opt *Options) {
		opt.CustomDefine[key] = value
	}
}

type Option func(*Options)
