package websockets

import (
	"fmt"
	// "io"
	"database/sql"
	"log"

	// "net/http"
	// db2 "real-time-forum/db"
	json2 "encoding/json"

	// "github.com/gorilla/websocket"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
)

func CreateMessageTable(db *sql.DB) {
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

func WsReadUsers(db *sql.DB) []byte {

	type Wsuser struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		LoginDate  string `json:"login_date"`
		LogoutDate string `json:"logout_date"`
		History	   string `json:"history"`
	}
	var data []Wsuser
	var json []byte
	var err error

	// Variables to use for assignment from database
	var id int
	var username string
	var loginDate string
	var logoutDate string
	var history string

	logger.InfoLogger.Println("GET: all users")

	// Select every row from user table
	rows, err := db.Query("SELECT id, username, login_date, logout_date, history FROM user WHERE id != 0 ORDER BY id")
	helper.CheckError(err)
	defer rows.Close()

	// log.Println("data1:", data)
	// Loop over every row
	for rows.Next() {

		rows.Scan(&id, &username, &loginDate, &logoutDate, &history)
		user := Wsuser{
			ID:         id,
			Username:   username,
			LoginDate:  loginDate,
			LogoutDate: logoutDate,
			History:	history,
		}

		data = append(data, user)
		// log.Println("data2:", data)

	}
	// log.Println("data3:", data)

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

func WsSaveMessage(db *sql.DB, body string, user_id string, target_id string, creation_time string) {

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

func WsReadMessages(db *sql.DB, messageUser string, messageTarget string) []byte {

	type Wsmessage struct {
		ID            int    `json:"id"`
		Body          string `json:"body"`
		User_id       string `json:"user_id"`
		Target_id     string `json:"target_id"`
		Creation_time string `json:"creation_time"`
	}

	returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)

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

	// log.Println("queryString:", queryString)

	rows, err := db.Query(queryString)
	helper.CheckError(err)
	defer rows.Close()

	// Loop over every row
	for rows.Next() {

		rows.Scan(&msgID, &msgBody, &msgUser, &msgTarget, &msgCreationTime)
		message := Wsmessage{
			ID:            msgID,
			Body:          msgBody,
			User_id:       msgUser,
			Target_id:     msgTarget,
			Creation_time: msgCreationTime,
		}

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

	returnedmessages = append(returnedmessages, json...)
	returnedmessages = append(returnedmessages, []byte(`}`)...)

	return returnedmessages

}

func WsSaveHistory(db *sql.DB, user int, target int) {
	// this function saves a message target into user's chat history array
	// has to save into both users' history

	// 1. query user history string from table

	// 2. turn data into array

	// 3. add target as first element of array

	// 4. turn array into string

	// 5. save string into user table history

}

// func WsReturnMessages() {
// 	returnedmessages := []byte(`{"type":"wsReturnedMessages","body":`)
// 	returnedmessages = append(returnedmessages, WsReadMessages(database, dat["user_id"].(string), dat["target_id"].(string))...)
// 	returnedmessages = append(returnedmessages, []byte(`}`)...)

// 	for client := range pool.Clients {

// 		// if received user Id conn is same as in Client struct, send messages back to this user
// 		if fmt.Sprintf("%d", client.UserID) == dat["user_id"].(string) {

// 			if err := client.Conn.WriteMessage(websocket.TextMessage, returnedmessages); err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}
// 	}
// }
