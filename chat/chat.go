package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var wsPool []*websocket.Conn

func onErr(err any) {
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/static/", handleAssets)
	http.HandleFunc("/ws", handleWs)

	port := ":9000"

	fmt.Println("Server is running on http://localhost" + port)

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
	wsPool = append(wsPool, ws)

	if err != nil {
		onErr(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			onErr(err)
			continue
		}

		for _, nWs := range wsPool {
			if err := nWs.WriteMessage(websocket.TextMessage, msg); err != nil {
				onErr(err)
			}
		}
	}
}
