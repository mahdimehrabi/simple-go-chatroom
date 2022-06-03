package websocket

import "github.com/gorilla/websocket"

//websocket connection wrapper
type WsConn struct {
	*websocket.Conn
}

type messageType string

var mtInfo messageType = "info"
var mtMessage messageType = "message"

type WsResponse struct {
	Action         string      `json:"action"`
	Username       string      `json:"username"`
	Message        interface{} `json:"message"`
	MessageType    messageType `json:"messageType"`
	Conn           *WsConn     `json:"-"`
	ConnectedUsers []string    `json:"connected_users"`
}

type Payload struct {
	Action      string      `json:"action"`
	Username    string      `json:"username"`
	Message     interface{} `json:"message"`
	MessageType messageType `json:"message"`
	Conn        *WsConn     `json:"-"`
}
