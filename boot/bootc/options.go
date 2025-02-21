package bootc

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/boot/logger"
)

// Options contains some configurations for current node
type Options struct {
	Initializer   channel.ChannelInitializer
	ChannelParams map[string]interface{}
	Logger        logger.Logger
}

var Default = &Options{
	ChannelParams: make(map[string]interface{}),
	Initializer:   nil,
}
