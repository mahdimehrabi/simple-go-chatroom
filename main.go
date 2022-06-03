package main

import (
	"chat/web"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	mux := web.Routes()
	log.Println("Starting websocket functionality 🚀")

	log.Println("Starting application 👄 on port:", port)
	http.ListenAndServe(port, mux)
}
