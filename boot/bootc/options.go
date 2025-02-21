package bootc

import (
	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/channel"
)

// Options contains some configurations for current node
type Options struct {
	Initializer   channel.ChannelInitializer
	ChannelParams map[string]interface{}
	Logger        boot.Logger
}

var Default = &Options{
	ChannelParams: make(map[string]interface{}),
	Initializer:   nil,
}
