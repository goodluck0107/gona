package channel

import (
	"errors"
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

type DefaultChannelHandlerContext struct {
	Name      string
	Next      *DefaultChannelHandlerContext
	Prev      *DefaultChannelHandlerContext
	mAttr     IAttr
	mPipeline ChannelPipeline
	mHandler  ChannelHandler
}

func newDefaultChannelHandlerContext(name string, pipeline ChannelPipeline, handler ChannelHandler, attr IAttr) (context *DefaultChannelHandlerContext) {
	if handler == nil {
		panic(errors.New("NullPointerException on NewDefaultEventHandlerContext handler is nil"))
	}
	context = new(DefaultChannelHandlerContext)
	context.Name = name
	context.mPipeline = pipeline
	context.mHandler = handler
	context.mAttr = attr
	return
}

func (this *DefaultChannelHandlerContext) handler() (handler ChannelHandler) {
	return this.mHandler
}

func (this *DefaultChannelHandlerContext) pipeline() (pipeline ChannelPipeline) {
	return this.mPipeline
}

func (this *DefaultChannelHandlerContext) channel() (handler Channel) {
	return this.pipeline().channel()
}

func (this *DefaultChannelHandlerContext) ID() string {
	return this.pipeline().channel().ID()
}

func (this *DefaultChannelHandlerContext) ContextAttr() IAttr {
	return this.mAttr
}

func (this *DefaultChannelHandlerContext) RemoteAddr() string {
	return this.pipeline().channel().RemoteAddr()
}

/*发起写事件，消息将被送往管道处理*/
func (this *DefaultChannelHandlerContext) Write(e interface{}) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	this.pipeline().FireMessageWrite(e)
}

/*发起关闭事件，消息将被送往管道处理*/
func (this *DefaultChannelHandlerContext) Close() {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	this.pipeline().FireChannelClose()
}

func (this *DefaultChannelHandlerContext) FireChannelActive() {
	//	fmt.Println(this.Name, "FireChannelActive")
	err, goonNext := this.invokeChannelActive()
	if err != nil {
		this.FireExceptionCaught(err)
	} else {
		if goonNext {
			if next := this.findContextInbound(); next != nil {
				next.FireChannelActive()
			}
		}
	}
}

func (this *DefaultChannelHandlerContext) FireMessageReceived(event interface{}) {
	//	fmt.Println(this.Name, "FireMessageReceived")
	ret, err, goonNext := this.invokeMessageReceived(event)
	if err != nil {
		this.FireExceptionCaught(err)
	} else {
		if goonNext {
			if next := this.findContextInbound(); next != nil {
				next.FireMessageReceived(ret)
			}
		}
	}
}

func (this *DefaultChannelHandlerContext) FireChannelInactive() {
	//	fmt.Println(this.Name, "FireChannelInactive")
	err, goonNext := this.invokeChannelInactive()
	if err != nil {
		this.FireExceptionCaught(err)
	} else {
		if goonNext {
			if next := this.findContextInbound(); next != nil {
				next.FireChannelInactive()
			}
		}
	}
}

func (this *DefaultChannelHandlerContext) FireExceptionCaught(err error) {
	//	fmt.Println(this.Name, "FireExceptionCaught")
	this.invokeExceptionCaught(err)
	if next := this.findContextInbound(); next != nil {
		next.FireExceptionCaught(err)
	}
}

func (this *DefaultChannelHandlerContext) FireMessageWrite(event interface{}) {
	//	fmt.Println(this.Name, "FireMessageWrite")
	ret, err := this.invokeWrite(event)
	if err != nil {
		this.FireExceptionCaught(err)
	} else {
		if next := this.findContextOutbound(); next != nil {
			next.FireMessageWrite(ret)
		} else {
			this.channel().Write(ret.([]byte))
		}
	}
}

func (this *DefaultChannelHandlerContext) FireChannelClose() {
	//	fmt.Println(this.Name, "FireChannelClose")
	err := this.invokeClose()
	if err != nil {
		this.FireExceptionCaught(err)
	} else {
		if next := this.findContextOutbound(); next != nil {
			next.FireChannelClose()
		} else {
			this.channel().Close()
		}
	}
}

func (this *DefaultChannelHandlerContext) invokeChannelActive() (err error, goonNext bool) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler, _ := this.handler().(ChannelInboundHandler)
	goonNext = handler.ChannelActive(this)
	return
}

func (this *DefaultChannelHandlerContext) invokeMessageReceived(event interface{}) (ret interface{}, err error, goonNext bool) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler, _ := this.handler().(ChannelInboundHandler)
	ret, goonNext = handler.MessageReceived(this, event)
	return
}

func (this *DefaultChannelHandlerContext) invokeChannelInactive() (err error, goonNext bool) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler, _ := this.handler().(ChannelInboundHandler)
	goonNext = handler.ChannelInactive(this)
	return
}

func (this *DefaultChannelHandlerContext) invokeExceptionCaught(err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler, _ := this.handler().(ChannelHandler)
	handler.ExceptionCaught(this, err)
}

func (this *DefaultChannelHandlerContext) invokeWrite(event interface{}) (ret interface{}, err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler, _ := this.handler().(ChannelOutboundHandler)
	ret = handler.Write(this, event)
	return
}

func (this *DefaultChannelHandlerContext) invokeClose() (err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler, _ := this.handler().(ChannelOutboundHandler)
	handler.Close(this)
	return
}

func (this *DefaultChannelHandlerContext) findContextInbound() *DefaultChannelHandlerContext {
	var ctx *DefaultChannelHandlerContext = this.Next
	for {
		isNil := (ctx == nil)
		if isNil {
			break
		}
		_, isType := ctx.handler().(ChannelInboundHandler)
		if isType {
			break
		}
		ctx = ctx.Next
	}
	return ctx
}

func (this *DefaultChannelHandlerContext) findContextOutbound() *DefaultChannelHandlerContext {
	var ctx *DefaultChannelHandlerContext = this.Prev
	for {
		isNil := (ctx == nil)
		if isNil {
			break
		}
		_, isType := ctx.handler().(ChannelOutboundHandler)
		if isType {
			break
		}
		ctx = ctx.Prev
	}
	return ctx
}
