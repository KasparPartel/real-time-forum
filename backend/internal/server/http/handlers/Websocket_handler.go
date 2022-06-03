package handlers

import (
	// "database/sql"
	// json2 "encoding/json"
	// "fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	// db2 "real-time-forum/db"
	// "real-time-forum/pkg/helper"
	// "real-time-forum/pkg/logger"
	"real-time-forum/pkg/websockets"
	// "time"
	// "real-time-forum/pkg/model"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin: func(r *http.Request) bool { return true }, // avoid CORS error
}



// func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
func WsEndpoint(w http.ResponseWriter, r *http.Request) /* (*websocket.Conn, error) */ {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // avoid CORS error

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		// return ws, err
	}
	log.Println("Backend: Client Successfully Connected to WebSocket...")
	log.Println("Backend: Attempting reader(ws)")

	// return ws, nil
	websockets.Reader(ws)
	go websockets.Writer(ws)

}



