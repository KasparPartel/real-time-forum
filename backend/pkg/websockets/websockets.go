package websockets

import (
	// "fmt"
	// "io"
	// "encoding/json"
	"log"
	"net/http"

	// db2 "real-time-forum/db"
	// "real-time-forum/pkg/helper"

	"github.com/gorilla/websocket"
	// "real-time-forum/pkg/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // avoid CORS error
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

// func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
//     ws, err := upgrader.Upgrade(w, r, nil)
//     if err != nil {
//         log.Println(err)
//         return ws, err
//     }
//     return ws, nil
// }

// func Reader(conn *websocket.Conn) {

// 	log.Println("Reader started")

// 	database, err := db2.OpenDB()
// 	helper.CheckError(err)
// 	CreateMessageTable(database)
// 	defer database.Close()

// 	type Message struct {

// 		// defining struct variables
// 		Type          string
// 		Body          string
// 		User_id       string
// 		Target_id     string
// 		Creation_time string
// 	}

// 	var incomingUser string
// 	var incomingTarget string

// 	for {
// 		messageType, p, err := conn.ReadMessage() // p == incoming message

// 		log.Println("Incoming message:", string(p))

// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		var incomingMessage Message

// 		err2 := json.Unmarshal(p, &incomingMessage)
// 		if err2 != nil {
// 			log.Println(err)
// 			return
// 		}

// 		log.Println("Hello there!")
// 		log.Println("incomingMessage:", incomingMessage)
// 		// log.Println("incomingMessage.Type: ", incomingMessage.Type)
// 		// log.Println("incomingMessage.Body: ", incomingMessage.Body)
// 		log.Println("incomingMessage.User_id: ", incomingMessage.User_id)
// 		log.Println("incomingMessage.Target_id: ", incomingMessage.Target_id)

// 		incomingUser = incomingMessage.User_id
// 		incomingTarget = incomingMessage.Target_id

// 		if incomingMessage.Type == "wsSaveChatMessage" {
// 			WsSaveMessage(
// 				database,
// 				incomingMessage.Body,
// 				incomingMessage.User_id,
// 				incomingMessage.Target_id,
// 				incomingMessage.Creation_time,
// 			)

// 			// time.Sleep(5 * time.Second)

// 			// log.Println("incomingUser:", incomingUser)
// 			// log.Println("incomingTarget:", incomingTarget)

// 			// this sends "Message saved" back to frontend
// 			if err := conn.WriteMessage(messageType, []byte(`{"type":"wsMessageSaved"}`)); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}

// 		if incomingMessage.Type == "wsGetChatMessages" && incomingMessage.User_id != "undefined" {
// 			// log.Println("incomingUser2:", incomingUser)
// 			// log.Println("incomingTarget2:", incomingTarget)

// 			returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
// 			returnedmessages = append(returnedmessages, WsReadMessages(database, incomingUser, incomingTarget)...)
// 			returnedmessages = append(returnedmessages, []byte(`}`)...)

// 			// log.Println("returnedmessages:", string(returnedmessages))

// 			// this send userlist from db back to frontend
// 			if err := conn.WriteMessage(messageType, returnedmessages); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}

// 		if incomingMessage.Type == "wsGetUsers" && incomingMessage.User_id != "undefined" {

// 			log.Println("Got wsGetUsers request from frontend")
// 			// log.Println(string(WsReadUsers(database)))

// 			// this the action with incoming message, rewrite to func -> db
// 			// log.Println("Printing out received message: ")
// 			// log.Println(string(p))

// 			// log.Println("readUsers(database):", string(WsReadUsers(database)))

// 			returnedusers := []byte(`{"type":"wsReturnedUsers","body":`)
// 			returnedusers = append(returnedusers, WsReadUsers(database)...)
// 			returnedusers = append(returnedusers, []byte(`}`)...)

// 			// log.Println("returnedusers:", string(returnedusers))

// 			// this send userlist from db back to frontend
// 			if err := conn.WriteMessage(messageType, returnedusers); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}

// 	}

// }

// func Writer(conn *websocket.Conn) {

// 	log.Println("Writer called")

	// for {
	// 	messageType, r, err := conn.NextReader()
	//     fmt.Println("Writer Sending:", r)
	//     if err != nil {
	// 		fmt.Println(err)
	//         return
	//     }
	//     w, err := conn.NextWriter(messageType)
	// 	// fmt.Println("Writer Sending:", string(w))
	//     if err != nil {
	//         fmt.Println(err)
	//         return
	//     }
	//     if _, err := io.Copy(w, r); err != nil {
	//         fmt.Println(err)
	//         return
	//     }
	//     if err := w.Close(); err != nil {
	//         fmt.Println(err)
	//         return
	//     }
	// }
//}
