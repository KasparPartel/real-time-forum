package handlers

import (
	"database/sql"

	"github.com/gorilla/websocket"

	// "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	db2 "real-time-forum/db"

	// config2 "real-time-forum/pkg/config"
	"real-time-forum/pkg/helper"
	// "real-time-forum/pkg/logger"
	// "real-time-forum/pkg/model"
	"encoding/json"
	"strconv"
)

// func connectDatabase() {

// 	database, err := db2.Open()
// 	helper.CheckError(err)
// 	createMessageTable(database)
// 	defer database.Close()

// }

// database, _ := sql.Open("sqlite3", "databsase.db")

// createTable(database)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {

	database, err := db2.Open()
	helper.CheckError(err)
	createMessageTable(database)
	defer database.Close()

	type Message struct {

		// defining struct variables
		Body          string
		User_id       string
		Target_id     string
		Creation_time string
	}

	for {
		messageType, p, err := conn.ReadMessage() // p == incoming message

		// defining a struct instance
		var incomingMessage Message

		// decoding incomingMessage struct
		// from json format
		err2 := json.Unmarshal(p, &incomingMessage)

		if err2 != nil {
			log.Println(err)
			return
		}

		// this the action with incoming message, rewrite to func -> db
		log.Println(string(p))
		// log.Println(p)
		// log.Println("Here?")
		saveMessage(
			database,
			incomingMessage.Body,
			incomingMessage.User_id,
			incomingMessage.Target_id,
			incomingMessage.Creation_time,
		)

		transmitMessageFromDB(database, incomingMessage.User_id)

		log.Println(tempMessages)
		for i := 0; i < len(tempMessages); i++ {
			log.Println(tempMessages[i])
		}

		// this repeats incoming message back to frontend
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

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

var tempMessages = []string{}

// func saveMessage(message string) {
// 	tempMessages = append(tempMessages, message)
//
// }

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

}

func transmitMessageFromDB(db *sql.DB, user_id string) {

	ID, _ := strconv.Atoi(user_id)
	rows, _ := db.Query(`SELECT body FROM messages WHERE user_id=?`, ID)

	var messageBody string

	for rows.Next() {
		rows.Scan(&messageBody)
		log.Println("Scanned from messages table: " + messageBody)
	}

}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // avoid CORS error

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Backend: Client Successfully Connected to WebSocket...")

	// connectDatabase()

	reader(ws)
}
