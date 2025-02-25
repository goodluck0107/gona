package bootc

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
)

// Options contains some configurations for current node
type Options struct {
	Initializer channel.ChannelInitializer
	Logger      logger.Logger
}

var Default = &Options{
	Initializer: nil,
}
