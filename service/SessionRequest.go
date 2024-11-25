package service

func NewSessionRequest(channelContext IChannelContext, reqContext IServiceContext) IServiceRequest {
	req := new(Request)
	req.mChannelContext = channelContext
	req.mReqContext = reqContext
	return req
}

type Request struct {
	mChannelContext IChannelContext
	mReqContext     IServiceContext
}

func (req *Request) ChannelContext() IChannelContext {
	return req.mChannelContext
}
func (req *Request) ReqContext() IServiceContext {
	return req.mReqContext
}
