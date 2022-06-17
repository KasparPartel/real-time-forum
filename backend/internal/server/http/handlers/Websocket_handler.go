package handlers

import (
	"fmt"
	"net/http"
	"real-time-forum/pkg/websockets"
)

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
}
