package boots

import "github.com/gorilla/websocket"

const (
	LocalHostIp              = "127.0.0.1"
	DefaultIp                = LocalHostIp
	DefaultPort       string = ":6829"
	DefaultClientPort        = 6829
	DefaultMsgType           = websocket.BinaryMessage
)
