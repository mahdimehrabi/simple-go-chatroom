package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WsConn struct {
	Conn *websocket.Conn
}

type WsResponse struct {
	Username    string `json:"username"`
	Message     string `json:"message"`
	MessageType string `json:"message"`
	Conn        WsConn
}

// func WsEndPoint(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgradeConnection.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("Client connected from %s", r.RemoteAddr)
// 	var response websocket
// }
