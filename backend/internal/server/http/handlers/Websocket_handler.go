package handlers

import (
	// "database/sql"
	// json2 "encoding/json"
	"fmt"
	// "github.com/gorilla/websocket"
	// "log"
	"net/http"
	// db2 "real-time-forum/db"
	// "real-time-forum/pkg/helper"
	// "real-time-forum/pkg/logger"
	"real-time-forum/pkg/websockets"
	// "time"
	// "real-time-forum/pkg/model"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	// CheckOrigin: func(r *http.Request) bool { return true }, // avoid CORS error
// }



// func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
// func WsEndpoint(w http.ResponseWriter, r *http.Request) {
func WsEndpoint(pool *websockets.Pool, w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("WebSocket Endpoint Hit")
    conn, err := websockets.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websockets.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()

	// ws, err := websockets.Upgrade(w, r)
    // if err != nil {
    //     fmt.Fprintf(w, "%+V\n", err)
    // }
    // go websockets.Writer(ws)
    // websockets.Reader(ws)

	// upgrader.CheckOrigin = func(r *http.Request) bool { return true } // avoid CORS error

	// ws, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	// return ws, err
	// }
	// log.Println("Backend: Client Successfully Connected to WebSocket...")
	// log.Println("Backend: Attempting reader(ws)")

	// // return ws, nil
	// go websockets.Writer(ws)
	// websockets.Reader(ws)

}



