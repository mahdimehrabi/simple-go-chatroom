package websocket

import "github.com/gorilla/websocket"

//websocket connection wrapper
type WsConn struct {
	*websocket.Conn
}

type messageType string

var mtInfo messageType = "info"
var mtMessage messageType = "message"

//the struct that will pass to user as response
type WsResponse struct {
	Action         string      `json:"action"`
	Username       string      `json:"username"`
	Message        interface{} `json:"message"`
	MessageType    messageType `json:"messageType"`
	Conn           *WsConn     `json:"-"`
	ConnectedUsers []string    `json:"connectedUsers"`
}

//the purpose of this struct is for receiving users data
type Payload struct {
	Action      string      `json:"action"`
	Username    string      `json:"username"`
	Message     interface{} `json:"message"`
	MessageType messageType `json:"messageType"`
	Conn        *WsConn     `json:"-"`
}
