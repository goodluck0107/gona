package service

func NewServiceResponse() *response {
	resp := new(response)
	return resp
}

type response struct {
	mMessage           interface{}
	mMessageSerializer ISerializer
}

func (resp *response) Write(msg interface{}, serializer ISerializer) {
	resp.mMessage = msg
	resp.mMessageSerializer = serializer
}
func (resp *response) Serialize() ([]byte, error) {
	return resp.mMessageSerializer.Serialize(resp.mMessage)
}
