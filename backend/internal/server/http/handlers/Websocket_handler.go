package handlers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage() // p == incoming message
		if err != nil {
			log.Println(err)
			return
		}

		// this the action with incoming message, rewrite to func -> db
		log.Println(string(p))
		// log.Println(p)
		// log.Println("Here?")
		saveMessage(string(p))

		log.Println(tempMessages)
		for i := 0; i < len(tempMessages); i++ {
			log.Println(tempMessages[i])
		}

		// this repeats incoming message back to frontend
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

}

var tempMessages = []string{}

func saveMessage(message string) {
	tempMessages = append(tempMessages, message)
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // avoid CORS error

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Backend: Client Successfully Connected to WebSocket...")

	reader(ws)
}
