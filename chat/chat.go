package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
	HandshakeTimeout: 0,
	ReadBufferSize: 0,
	WriteBufferSize: 0,
	WriteBufferPool: nil,
	Subprotocols: []string{},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		panic("TODO")
	},
	EnableCompression: false,
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/static/", handleAssets)
	http.HandleFunc("/ws/", handleWs)

	port := ":9000"

	fmt.Println("Server is running on port " + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func handleAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.RequestURI[1:])
}

func handleWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic("Something went wrong")
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			panic("Something went wrong")
		}

		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			panic("Something went wrong")
		}
	}
}
