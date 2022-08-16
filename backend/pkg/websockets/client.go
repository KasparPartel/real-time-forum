package websockets

import (
	"fmt"
	"log"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"strconv"

	// "sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	UserID int
	Conn   *websocket.Conn
	Pool   *Pool
}

// check how this conforms with previous message types!!!
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
	// Client *string `json:"client"`
	Conn *websocket.Conn `json:"conn"`
}

func (c *Client) Read() {

	database, err := db2.OpenDB()
	helper.CheckError(err)
	// CreateMessageTable(database)
	defer database.Close()

	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
		// sends users update on conn status closing
		WsSendUsers(database, strconv.Itoa(c.UserID), c.Pool)
	}()

	log.Println("c Client Read()")
	// connection := c.Conn.UnderlyingConn()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p), Conn: c.Conn}
		c.Pool.Broadcast <- message
		fmt.Printf("POOL: Message Received: %+v\n", message)
	}
}
