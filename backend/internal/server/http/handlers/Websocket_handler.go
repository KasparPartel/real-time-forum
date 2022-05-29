package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	// "time"
	// "real-time-forum/pkg/model"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {

	log.Println("Reader started")

	database, err := db2.Open()
	helper.CheckError(err)
	createMessageTable(database)
	defer database.Close()

	type Message struct {

		// defining struct variables
		Type          string
		Body          string
		User_id       string
		Target_id     string
		Creation_time string
	}

	var incomingUser string
	var incomingTarget string

	for {
		messageType, p, err := conn.ReadMessage() // p == incoming message

		if err != nil {
			log.Println(err)
			return
		}

		var incomingMessage Message

		err2 := json2.Unmarshal(p, &incomingMessage)
		if err2 != nil {
			log.Println(err)
			return
		}

		log.Println(incomingMessage)
		log.Println("incomingMessage.Type: ", incomingMessage.Type)
		log.Println("incomingMessage.Body: ", incomingMessage.Body)
		log.Println("incomingMessage.User_id: ", incomingMessage.User_id)
		log.Println("incomingMessage.Target_id: ", incomingMessage.Target_id)

		incomingUser = incomingMessage.User_id
		incomingTarget = incomingMessage.Target_id

		if incomingMessage.Type == "wsSaveChatMessage" {
			saveMessage(
				database,
				incomingMessage.Body,
				incomingMessage.User_id,
				incomingMessage.Target_id,
				incomingMessage.Creation_time,
			)

			// time.Sleep(5 * time.Second)

			log.Println("incomingUser:", incomingUser)
			log.Println("incomingTarget:", incomingTarget)

			// this send "Message saved" back to frontend
			// if err := conn.WriteMessage(messageType, []byte(`"type":"wsMessageSaved"`)); err != nil {
			if err := conn.WriteMessage(messageType, []byte(`{"type":"wsMessageSaved"}`)); err != nil {
				log.Println(err)
				return
			}

			// returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
			// returnedmessages = append(returnedmessages, readMessages(database, incomingMessage.User_id, incomingMessage.Target_id)...)
			// returnedmessages = append(returnedmessages, []byte(`}`)...)

			// log.Println("returnedmessages:", string(returnedmessages))

			// // this send userlist from db back to frontend
			// if err := conn.WriteMessage(messageType, returnedmessages); err != nil {
			// 	log.Println(err)
			// 	return
			// }
		}

		if incomingMessage.Type == "wsGetChatMessages" {
			log.Println("incomingUser2:", incomingUser)
			log.Println("incomingTarget2:", incomingTarget)

			returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
			returnedmessages = append(returnedmessages, readMessages(database, incomingUser, incomingTarget)...)
			returnedmessages = append(returnedmessages, []byte(`}`)...)

			log.Println("returnedmessages:", string(returnedmessages))

			// this send userlist from db back to frontend
			if err := conn.WriteMessage(messageType, returnedmessages); err != nil {
				log.Println(err)
				return
			}
		}

		if incomingMessage.Type == "wsGetUsers" {

			log.Println("Got wsGetUsers request from frontend")
			log.Println(string(readUsers(database)))

			// this the action with incoming message, rewrite to func -> db
			log.Println("Printing out received message: ")
			log.Println(string(p))

			returnedusers := []byte(`{"type":"wsReturnedUsers","body":`)
			returnedusers = append(returnedusers, readUsers(database)...)
			returnedusers = append(returnedusers, []byte(`}`)...)

			log.Println("returnedusers:", string(returnedusers))

			// this send userlist from db back to frontend
			if err := conn.WriteMessage(messageType, returnedusers); err != nil {
				log.Println(err)
				return
			}
		}

	}

}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // avoid CORS error

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Backend: Client Successfully Connected to WebSocket...")
	log.Println("Backend: Attempting reader(ws)")

	reader(ws)
}

func createMessageTable(db *sql.DB) {
	messages_table := `CREATE TABLE IF NOT EXISTS messages (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "body" TEXT,
        "user_id" TEXT,
        "target_id" TEXT,
        "creation_time" TEXT);`
	query, err := db.Prepare(messages_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	log.Println("Messages Table created successfully!")
}

func readUsers(db *sql.DB) []byte {

	type Wsuser struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		LoginTime string `json:"login_Time"`
	}
	var data []Wsuser
	var json []byte
	var err error

	// Variables to use for assignment from database
	var userID int
	var username string
	var loginDate string

	logger.InfoLogger.Println("GET: all users")

	// Select every row from user table
	// rows, err := db.Query("SELECT user_id, username, login_date FROM user WHERE user_id=?", 1)
	// rows, err := db.Query("SELECT user_id, username, login_date FROM user ORDER BY user_id LIMIT -1 OFFSET 1")
	rows, err := db.Query("SELECT user_id, username, login_date FROM user ORDER BY user_id")
	// rows, err := db.Query("SELECT * FROM user")
	helper.CheckError(err)
	defer rows.Close()

	// Loop over every row
	for rows.Next() {

		user := Wsuser{
			ID:        userID,
			Username:  username,
			LoginTime: loginDate,
		}
		rows.Scan(&userID, &username, &loginDate)

		data = append(data, user)
	}

	if len(data) == 0 {
		logger.WarningLogger.Println("There are 0 users")
	}

	// Write json to return
	json, err = json2.Marshal(data)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return json

}

func saveMessage(db *sql.DB, body string, user_id string, target_id string, creation_time string) {

	message := `INSERT INTO messages(body, user_id, target_id, creation_time) VALUES (?, ?, ?, ?)`
	query, err := db.Prepare(message)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(body, user_id, target_id, creation_time)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved message to db: ", body)
}

func readMessages(db *sql.DB, messageUser string, messageTarget string) []byte {

	type Wsmessage struct {
		ID            int    `json:"id"`
		Body          string `json:"body"`
		User_id       string `json:"user_id"`
		Target_id     string `json:"target_id"`
		Creation_time string `json:"creation_time"`
	}

	var data []Wsmessage
	var json []byte
	var err error

	// Variables to use for assignment from database
	var msgID int
	var msgBody string
	var msgUser string
	var msgTarget string
	var msgCreationTime string

	logger.InfoLogger.Println("GET: all messages with current user and target")

	queryString := fmt.Sprintf("%s%s%s%s%s%s%s%s",
		"SELECT * from messages WHERE user_id=",
		messageUser,
		" AND target_id=",
		messageTarget,
		" OR user_id=",
		messageTarget,
		" AND target_id=",
		messageUser)

	log.Println("queryString:", queryString)

	rows, err := db.Query(queryString)
	helper.CheckError(err)
	defer rows.Close()

	// Loop over every row
	for rows.Next() {

		message := Wsmessage{
			ID:            msgID,
			Body:          msgBody,
			User_id:       msgUser,
			Target_id:     msgTarget,
			Creation_time: msgCreationTime,
		}
		rows.Scan(&msgID, &msgBody, &msgUser, &msgTarget, &msgCreationTime)

		data = append(data, message)
	}

	if len(data) == 0 {
		logger.WarningLogger.Println("There are 0 corresponding messages")
	}

	// Write json to return
	json, err = json2.Marshal(data)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return json

}
