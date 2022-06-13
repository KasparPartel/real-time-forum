package websockets

import (
    "fmt"
    // "io"
    "log"
	"database/sql"
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
		ID        int    `json:"id"`
		Username  string `json:"username"`
		LoginDate string `json:"login_date"`
		LogoutDate string `json:"logout_date"`
	}
	var data []Wsuser
	var json []byte
	var err error

	// Variables to use for assignment from database
	var id int
	var username string
	var loginDate string
	var logoutDate string

	logger.InfoLogger.Println("GET: all users")

	// Select every row from user table
	rows, err := db.Query("SELECT id, username, login_date, logout_date FROM user WHERE id != 0 ORDER BY id")
	helper.CheckError(err)
	defer rows.Close()

	// log.Println("data1:", data)
	// Loop over every row
	for rows.Next() {
		
		rows.Scan(&id, &username, &loginDate, &logoutDate)
		user := Wsuser{
			ID:        id,
			Username:  username,
			LoginDate: loginDate,
			LogoutDate: logoutDate,
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

	return json

}
