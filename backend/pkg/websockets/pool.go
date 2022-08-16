package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"strconv"

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
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			for user := range pool.Clients {
				fmt.Println(client)
				user.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})

				if user.ID == client.ID {
					user.Conn.WriteJSON(Message{Type: 1, Body: `"newClient":` + string(client.ID)})
				}

			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("POOL: Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				// for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}

		case message := <-pool.Broadcast:
			byt := []byte(message.Body)
			var dat map[string]interface{}
			if err := json.Unmarshal(byt, &dat); err != nil {
				panic(err)
			}

			// here we separate incoming messages by type
			// if the frontend sends user ID for this ws conn
			if dat["type"] == "sendUser" {

				for client := range pool.Clients {

					// if received user Id conn is same as in Client struct, save user ID in Client
					if client.Conn == message.Conn {

						if dat["activeUser"] != nil {

							client.UserID = int(dat["activeUser"].(float64))
							fmt.Println("Active user received and saved to client.UserID:", client.UserID)

							for client := range pool.Clients {
								fmt.Printf("Active client UserID in pool: %d\n", client.UserID)
							}
						}
					}
				}
			}

			// if frontend sends Modal clicked event, last seen history count is saved into db
			if dat["type"] == "sendModal" {
				_, history := getHistory(database, int(dat["activeUser"].(float64)))
				updateHistory(database, history, int(dat["activeUser"].(float64)), int(dat["targetUser"].(float64)), 0)

				// this send userlist from db back to frontend that sent request
				for client := range pool.Clients {
					// if received user Id conn is same as in Client struct, send users back to this user
					if fmt.Sprintf("%d", client.UserID) == strconv.Itoa(int(dat["activeUser"].(float64))) {
						modalString := fmt.Sprintf(`{"type":"wsModalSaved","user":"%d","target":"%d"}`,
							int(dat["activeUser"].(float64)),
							int(dat["targetUser"].(float64)))

						if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(modalString)); err != nil {
							log.Println(err)
							return
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

				// this sends "Message saved" back to frontend connection
				for client := range pool.Clients {

					// if received user Id conn is same as in Client struct, send confirmation back
					if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) ||
						fmt.Sprintf("%d", client.UserID) == dat["target_id"].(string) {

						if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"wsMessageSaved"}`)); err != nil {
							log.Println(err)
							return
						}
					}
				}
			}
			// sends users to all frontends when asked
			if dat["type"] == "wsGetUsers" && dat["activeUser"] != "undefined" /* && dat["target_id"] != "undefined" */ {
				WsSendUsers(database, strconv.Itoa(int(dat["activeUser"].(float64))), pool)
			}
			// sends users to all frontends on user login/logout status change
			if dat["type"] == "wsUserStatusChange" {
				WsSendUsers(database, strconv.Itoa(int(dat["changedUser"].(float64))), pool)
			}

			if dat["type"] == "wsGetChatMessages" && dat["user_id"] != "undefined" && dat["target_id"] != "undefined" {

				jsonmessages, _ := WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string), dat["count"].(string))

				returnedmessages := []byte(fmt.Sprintf(`{"type":"wsReturnedMessages","user":"%s","target":"%s","body":`, dat["user_id"].(string), dat["target_id"].(string)))
				returnedmessages = append(returnedmessages, jsonmessages...)
				returnedmessages = append(returnedmessages, []byte(`}`)...)

				for client := range pool.Clients {
					// if received user Id conn is same as in Client struct, send messages back to this user
					if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) ||
						fmt.Sprintf("%d", client.UserID) == dat["target_id"].(string) {
						if err := client.Conn.WriteMessage(websocket.TextMessage, returnedmessages); err != nil {
							log.Println(err)
							return
						}
					}
				}
			}
		}
	}
}
