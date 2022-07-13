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

					// if received user Id conn is same as in Client struct, save user ID in Client
					if client.Conn == message.Conn {

						if dat["activeUser"] != nil { // trying to fix bug where "sendUser" has no "activeUser"

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
				log.Println("Received sendModal:", dat)
				user := strconv.Itoa(int(dat["activeUser"].(float64)))
				target := strconv.Itoa(int(dat["targetUser"].(float64)))
				_, dbMessages := WsReadMessages(database, user, target)
				_, history := getHistory(database, int(dat["activeUser"].(float64)))
				updateHistory(database, history, int(dat["activeUser"].(float64)), int(dat["targetUser"].(float64)), dbMessages)
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

			if dat["type"] == "wsGetUsers" && dat["activeUser"] != "undefined" /* && dat["target_id"] != "undefined" */ {

				// log.Println("Got wsGetUsers request from frontend")
				// unreadArray := []int{}
				// _, userArray := WsReadUsers(database)
				// user := strconv.Itoa(int(dat["activeUser"].(float64)))
				// _, history := getHistory(database, int(dat["activeUser"].(float64)))

				// for i := 0; i < len(userArray); i++ {
				// 	if userArray[i] != int(dat["activeUser"].(float64)) {
				// 		target := strconv.Itoa(userArray[i])
				// 		_, dbMsgLength := WsReadMessages(database, user, target)
				// 		// _, historyMsgLength := WsReadMessages(database, user, target)
				// 		if compareHistory(history, userArray[i], dbMsgLength) {
				// 			unreadArray = append(unreadArray, userArray[i])
				// 		}
				// 	}
				// }
				// log.Println("Unread array:", unreadArray)

				// userpool := []byte(`,"pool":"`)
				// for client := range pool.Clients {
				// 	// userpool = append(userpool, client.UserID)
				// 	userpool = append(userpool, []byte(strconv.Itoa(client.UserID))...)
				// 	userpool = append(userpool, []byte(`,`)...)
				// }
				// userpool = userpool[:len(userpool)-1]
				// userpool = append(userpool, []byte(`"`)...)

				// fmt.Println("pool.Clients")
				// fmt.Println(pool.Clients)

				// unreadpool := []byte(`,"unread":"`)
				// if len(unreadArray) > 0 {
				// 	for user := range unreadArray {
				// 		unreadpool = append(unreadpool, []byte(strconv.Itoa(user))...)
				// 		unreadpool = append(unreadpool, []byte(`,`)...)
				// 	}
				// 	unreadpool = unreadpool[:len(unreadpool)-1]
				// }
				// unreadpool = append(unreadpool, []byte(`"`)...)

				// userJson, _ := WsReadUsers(database)

				// returnedusers := []byte(`{"type":"wsReturnedUsers","body":`)
				// returnedusers = append(returnedusers, userJson...)
				// returnedusers = append(returnedusers, userpool...)
				// returnedusers = append(returnedusers, unreadpool...)
				// returnedusers = append(returnedusers, []byte(`}`)...)

				// var returnedusers []byte
				returnedusers := WsReturnUsers(database, strconv.Itoa(int(dat["activeUser"].(float64))), pool)

				// this send userlist from db back to frontend that sent request
				for client := range pool.Clients {

					// SEND HERE ONLY TO REQUESTING USER

					// if received user Id conn is same as in Client struct, send users back to this user
					//if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) {

					if err := client.Conn.WriteMessage(websocket.TextMessage, returnedusers); err != nil {
						log.Println(err)
						return
					}
					//}
				}
			}

			if dat["type"] == "wsGetChatMessages" && dat["user_id"] != "undefined" && dat["target_id"] != "undefined" {
				// returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
				// returnedmessages = append(returnedmessages, WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string))...)
				// returnedmessages = append(returnedmessages, []byte(`}`)...)
				log.Println("Got wsGetChatMessages: user_id, target_id", dat["user_id"], dat["target_id"])

				// type Wsmessage struct {
				// 	ID            int    `json:"id"`
				// 	Body          string `json:"body"`
				// 	User_id       string `json:"user_id"`
				// 	Target_id     string `json:"target_id"`
				// 	Creation_time string `json:"creation_time"`
				// }

				jsonmessages, messagelength := WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string))

				log.Println("Nr of returnedmessages:", messagelength)

				// Write json to return
				// json, err := json2.Marshal(returndata)
				// if err != nil {
				// 	logger.ErrorLogger.Println(err)
				// }

				// log.Println("WsReadMessages length:", len(data))
				returnedmessages := []byte(fmt.Sprintf(`{"type":"wsReturnedMessages","sender":"%s","body":`, dat["user_id"].(string)))
				returnedmessages = append(returnedmessages, jsonmessages...)
				returnedmessages = append(returnedmessages, []byte(`}`)...)

				log.Println("pool.Clients", pool.Clients)

				for client, _ := range pool.Clients {
					log.Println("User in Pool client.UserID:", client.UserID)
					// if received user Id conn is same as in Client struct, send messages back to this user
					if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) ||
						fmt.Sprintf("%d", client.UserID) == dat["target_id"].(string) { // PROBLEM?
						log.Println("Sending messages to user:", client.UserID) // or target_id?
						// log.Println("Sending messages:", string(returnmessages))
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
