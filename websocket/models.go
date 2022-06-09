package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
)

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

//readJson return action,payloadStruct,erro
func (wc *WsConn) ReadJSON() (string, any, error) {
	_, r, err := wc.NextReader()
	if err != nil {
		return "", nil, err
	}
	var payload Payload
	err = ioReaderJson(r, &payload)
	if err != nil {
		return "", nil, err
	}
	switch payload.Action {
	case "message":
		var payloadMessage PayloadMessage
		err = ioReaderJson(r, &payloadMessage)
		if err != nil {
			return "", nil, err
		}
		return "message", payloadMessage, err
	}
	return "", nil, err
}

func ioReaderJson(r io.Reader, payload any) error {
	err := json.NewDecoder(r).Decode(payload)
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}
