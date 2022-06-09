package websocket

import "github.com/gorilla/websocket"

//the struct that will pass to user as response
type WsResponse struct {
	Action         string      `json:"action"`
	Username       string      `json:"username"`
	Message        interface{} `json:"message"`
	Conn           *WsConn     `json:"-"`
	ConnectedUsers []string    `json:"connectedUsers"`
}

//the purpose of this struct is for receiving users data
type Payload struct {
	Action   string  `json:"action"`
	Username string  `json:"username"`
	Conn     *WsConn `json:"-"`
}

type PayloadMessage struct {
	Payload
	Message interface{} `json:"message"`
}

//websocket connection wrapper
type WsConn struct {
	*websocket.Conn
}
