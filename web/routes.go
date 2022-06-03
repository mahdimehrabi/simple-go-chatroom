package web

import (
	"net/http"

	"chat/websocket"

	"github.com/bmizerany/pat"
)

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(Home))
	mux.Get("/ws", http.HandlerFunc(websocket.WsEndPoint))

	go websocket.ListenToWsChannel()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
