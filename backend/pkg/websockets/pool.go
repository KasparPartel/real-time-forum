package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"

	"github.com/gorilla/websocket"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {

	database, err := db2.OpenDB()
	helper.CheckError(err)
	CreateMessageTable(database)
	defer database.Close()

	for {
		// var activeClientID string
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("POOL: Size of Connection Pool: ", len(pool.Clients))
			fmt.Println("pool.Clients", pool.Clients)
			fmt.Println("connected client ID:", client.ID)
			// activeClientID = client.ID

			for user, _ := range pool.Clients {
				// for client := range pool.Clients {
				fmt.Println(client)
				user.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})

				if user.ID == client.ID {
					user.Conn.WriteJSON(Message{Type: 1, Body: `"newClient":` + string(client.ID)})
				}

			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("POOL: Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				// for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("POOL: Sending message to all clients in Pool")
			// fmt.Println("POOL: incoming message:", message)
			// fmt.Println("POOL: incoming message.body:", message.Body)
			// fmt.Println("POOL: incoming message.client:", message.Conn)

			byt := []byte(message.Body)
			fmt.Println("Trying to unmarshal")
			fmt.Println("message.Body", message.Body)
			// fmt.Println("message.Body byt", byt)
			var dat map[string]interface{}
			if err := json.Unmarshal(byt, &dat); err != nil {
				panic(err)
			}
			fmt.Println("Unmarshaled data:", dat)

			// returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
			// returnedmessages = append(returnedmessages, WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string))...)
			// returnedmessages = append(returnedmessages, []byte(`}`)...)

			// here we separate incoming messages by type
			// if the frontend sends user ID for this ws conn
			if dat["type"] == "sendUser" {

				for client, _ := range pool.Clients {

					// if recerived user Id conn is same as in Client struct, save user ID in Client
					if client.Conn == message.Conn {

						client.UserID = int(dat["activeUser"].(float64))
						fmt.Println("Active user received and saved to client.UserID:", client.UserID)

						for client := range pool.Clients {
							fmt.Printf("Active client UserID in pool: %d\n", client.UserID)
						}
					}
				}
			}

			// if the frontend sends Message to be saved into db
			if dat["type"] == "wsSaveChatMessage" {

				WsSaveMessage(
					database,
					dat["body"].(string),
					dat["user_id"].(string),
					dat["target_id"].(string),
					dat["creation_time"].(string),
				)
				WsSaveHistory(database, dat["user_id"].(string), dat["target_id"].(string))
				WsSaveHistory(database, dat["target_id"].(string), dat["user_id"].(string))

				// this sends "Message saved" back to frontend connection
				for client := range pool.Clients {

					// if received user Id conn is same as in Client struct, send confirmation back
					if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) {

						if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"wsMessageSaved"}`)); err != nil {
							log.Println(err)
							return
						}
					}
				}
			}

			if dat["type"] == "wsGetUsers" && dat["user_id"] != "undefined" {

				log.Println("Got wsGetUsers request from frontend")

				returnedusers := []byte(`{"type":"wsReturnedUsers","body":`)
				returnedusers = append(returnedusers, WsReadUsers(database)...)
				returnedusers = append(returnedusers, []byte(`}`)...)

				// this send userlist from db back to frontend that sent request
				for client := range pool.Clients {

					// if received user Id conn is same as in Client struct, send users back to this user
					//if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) {

					if err := client.Conn.WriteMessage(websocket.TextMessage, returnedusers); err != nil {
						log.Println(err)
						return
					}
					//}
				}
			}

			if dat["type"] == "wsGetChatMessages" && dat["user_id"] != "undefined" {
				// returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
				// returnedmessages = append(returnedmessages, WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string))...)
				// returnedmessages = append(returnedmessages, []byte(`}`)...)
				log.Println("Got wsGetChatMessages: user_id, target_id", dat["user_id"], dat["target_id"])

				returnmessages := WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string))

				// log.Println("returnessages:", string(returnmessages))

				log.Println("pool.Clients", pool.Clients)

				for client, _ := range pool.Clients {
					log.Println("User in Pool client.UserID:", client.UserID)
					// if received user Id conn is same as in Client struct, send messages back to this user
					if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) ||
						fmt.Sprintf("%d", client.UserID) == dat["target_id"].(string) { // PROBLEM?
						log.Println("Sending messages to user:", dat["user_id"])
						// log.Println("Sending messages:", string(returnmessages))
						if err := client.Conn.WriteMessage(websocket.TextMessage, returnmessages); err != nil {
							log.Println(err)
							return
						}
					}
				}
			}
		}
	}
}
