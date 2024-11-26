package channel

import (
	"sync"
)

type TailHandler struct {
}

func (this *TailHandler) Write(ctx ChannelContext, e interface{}) (ret interface{}) {
	return e
}

func (this *TailHandler) Close(ctx ChannelContext) {
}

func (this *TailHandler) ExceptionCaught(ctx ChannelContext, err error) {
}

type HeadHandler struct {
}

func (this *HeadHandler) MessageReceived(ctx ChannelContext, e interface{}) (ret interface{}, goonNext bool) {
	return e, true
}

func (this *HeadHandler) ChannelActive(ctx ChannelContext) (goonNext bool) {
	return true
}

func (this *HeadHandler) ChannelInactive(ctx ChannelContext) (goonNext bool) {
	return true
}

func (this *HeadHandler) ExceptionCaught(ctx ChannelContext, err error) {
}

type DefaultChannelPipeline struct {
	Lock     *sync.Mutex
	Head     *DefaultChannelHandlerContext
	Tail     *DefaultChannelHandlerContext
	mChannel Channel
	mAttr    IAttr
}

func NewDefaultChannelPipeline(channel Channel) (pipeline *DefaultChannelPipeline) {
	pipeline = new(DefaultChannelPipeline)
	pipeline.mChannel = channel
	pipeline.mAttr = channel
	pipeline.Lock = new(sync.Mutex)

	tailHandler := new(TailHandler)
	pipeline.Tail = newDefaultChannelHandlerContext("Tail", pipeline, tailHandler, pipeline.mAttr)

	headHandler := new(HeadHandler)
	pipeline.Head = newDefaultChannelHandlerContext("Head", pipeline, headHandler, pipeline.mAttr)

	pipeline.Head.Next = pipeline.Tail
	pipeline.Tail.Prev = pipeline.Head
	return
}

func (this *DefaultChannelPipeline) channel() (channel Channel) {
	return this.mChannel
}

func (this *DefaultChannelPipeline) ContextAttr() IAttr {
	return this.mAttr
}

func (pipeline *DefaultChannelPipeline) AddFirst(name string, handler ChannelHandler) ChannelPipeline {
	pipeline.Lock.Lock()
	defer pipeline.Lock.Unlock()

	newCtx := newDefaultChannelHandlerContext(name, pipeline, handler, pipeline.ContextAttr())

	nextCtx := pipeline.Head.Next
	newCtx.Prev = pipeline.Head
	newCtx.Next = nextCtx
	pipeline.Head.Next = newCtx
	nextCtx.Prev = newCtx
	return pipeline
}

func (pipeline *DefaultChannelPipeline) AddLast(name string, handler ChannelHandler) ChannelPipeline {
	pipeline.Lock.Lock()
	defer pipeline.Lock.Unlock()

	newCtx := newDefaultChannelHandlerContext(name, pipeline, handler, pipeline.ContextAttr())

	prev := pipeline.Tail.Prev
	newCtx.Prev = prev
	newCtx.Next = pipeline.Tail
	prev.Next = newCtx
	pipeline.Tail.Prev = newCtx

	return pipeline
}

func (this *DefaultChannelPipeline) FireChannelActive() {
	//	fmt.Println("DefaultEventPipeline fireChannelActive")
	this.Head.FireChannelActive()
}

func (this *DefaultChannelPipeline) FireMessageReceived(event interface{}) {
	//	fmt.Println("DefaultEventPipeline FireMessageReceived")
	this.Head.FireMessageReceived(event)
}

func (this *DefaultChannelPipeline) FireChannelInactive() {
	//	fmt.Println("DefaultEventPipeline FireChannelInactive")
	this.Head.FireChannelInactive()
}

func (this *DefaultChannelPipeline) FireExceptionCaught(err error) {
	//	fmt.Println("DefaultEventPipeline FireExceptionCaught")
	this.Head.FireExceptionCaught(err)
}

func (this *DefaultChannelPipeline) FireMessageWrite(event interface{}) {
	//	fmt.Println("DefaultEventPipeline FireMessageWrite")
	this.Tail.FireMessageWrite(event)
}

func (this *DefaultChannelPipeline) FireChannelClose() {
	//	fmt.Println("DefaultEventPipeline FireChannelClose")
	this.Tail.FireChannelClose()
}
