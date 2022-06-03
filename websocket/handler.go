package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[WsConn]string)
var wsChan = make(chan Payload)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected from %s", r.RemoteAddr)
	var response WsResponse
	response.Action = "wellcome"
	response.Message = "<em><small>Wellcome to chat</small></em>"

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	conn := WsConn{Conn: ws}
	clients[conn] = ""

	go ListenForWs(&conn)

}

func ListenForWs(conn *WsConn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()
	for {
		var payload Payload
		err := conn.ReadJSON(&payload)
		if err != nil {

		} else {
			payload.Conn = conn
			wsChan <- payload
		}
	}

}

func ListenToWsChannel() {
	var response WsResponse
	for {
		e := <-wsChan
		switch e.Action {
		case "connected":
			if e.Username != "" {
				response.Action = "connected"
				response.Username = e.Username
				response.MessageType = mtInfo
				response.Message = e.Username + " Connected!"
			}

		}
	}
}

func BroadCastToAll(response WsResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Printf("Websocket error on %s: %s", response.Action, err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}
