package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[WsConn]string)
var wsChan = make(chan Payload)

//connection upgrader
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//main end point for websocket
//this is root of handling all websocket requests
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

//listen for receving client requests(payload) as json
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

//handling user requests
func ListenToWsChannel() {
	var response WsResponse
	for {
		e := <-wsChan
		switch e.Action {
		case "username":
			if e.Username != "" {
				response.Action = "username"
				response.Username = e.Username
				response.MessageType = mtInfo
				if clients[*e.Conn] == "" {
					response.Message = e.Username + " Connected!"
				} else {
					fmt.Println(clients[*e.Conn])
					response.Message = clients[*e.Conn] + " changed its name to " + e.Username
				}
				clients[*e.Conn] = e.Username
				response.ConnectedUsers = getConnectedUsers(clients)
				broadCastAllExcept(response, *e.Conn)

				response.Action = "connectedUsers"
				response.Message = ""
				err := e.Conn.WriteJSON(response)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "message":
			response.Action = "message"
			response.Message = e.Message.(string)
			response.MessageType = "message"
			broadCastAll(response)
		}
	}
}

//broadcast a data to all clients except exceptConn
func broadCastAllExcept(response WsResponse, excpetConn WsConn) {
	for client := range clients {
		if client == excpetConn {
			continue
		}
		err := client.WriteJSON(response)
		if err != nil {
			log.Printf("Websocket error on %s: %s", response.Action, err)
			LeftRoom(&client)
			if response.Action == "connectedUsers" {
				return
			}
		}
	}
}

//broadcast a data to all clients
func broadCastAll(response WsResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Printf("Websocket error on %s: %s", response.Action, err)
			LeftRoom(&client)
		}
	}
}

func LeftRoom(client *WsConn) {
	_ = client.Close()
	message := clients[*client] + " left the room!"
	delete(clients, *client)
	connectedResponse := WsResponse{
		Action:         "left",
		Username:       "mahd",
		Message:        message,
		MessageType:    "info",
		ConnectedUsers: getConnectedUsers(clients),
	}

	broadCastAll(connectedResponse)
}

//collect connected user names as string array
func getConnectedUsers(clients map[WsConn]string) []string {
	users := make([]string, 0)
	for _, v := range clients {
		if v != "" {
			users = append(users, v)
		}
	}
	return users
}
