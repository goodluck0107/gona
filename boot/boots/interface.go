package boots

import (
	"fmt"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
)

func Serve(opts ...Option) {
	opt := Default
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

func WithChannelParams(channelParams map[string]interface{}) Option {
	return func(opt *Options) {
		opt.ChannelParams = channelParams
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

type Option func(*Options)

func check(opts *Options) error {
	if opts.Initializer == nil {
		return fmt.Errorf("channel initializer is nil")
	}
	if opts.ChannelParams == nil {
		opts.ChannelParams = make(map[string]interface{})
	}
	if opts.ReadTimeOut != 0 {
		opts.ChannelParams[channel.KeyReadTimeOut] = opts.ReadTimeOut
	}
	if opts.WriteTimeOut != 0 {
		opts.ChannelParams[channel.KeyWriteTimeOut] = opts.WriteTimeOut
	}
	return nil
}
