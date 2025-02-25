package bootc

import (
	"fmt"

	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
	"github.com/mohae/deepcopy"
)

func Serv(opts ...Option) listener.IConnector {
	opt := deepcopy.Copy(Default).(*Options)
	for _, option := range opts {
		option(opt)
	}
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
	n := &ClientBootStrap{
		Options: opt,
	}
	n.connector, n.acceptor = listener.Create()
	startupE := n.startup()
	if startupE != nil {
		panic(startupE)
	}
	return n.connector
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

type Option func(*Options)

func check(opts *Options) error {
	if opts.Initializer == nil {
		return fmt.Errorf("channel initializer is nil")
	}
	return nil
}
