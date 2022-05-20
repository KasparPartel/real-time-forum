package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
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

		// if incomingMessage.Type == "wsSaveChatMessage" {
		// 	saveMessage(
		// 		database,
		// 		incomingMessage.Body,
		// 		incomingMessage.User_id,
		// 		incomingMessage.Target_id,
		// 		incomingMessage.Creation_time,
		// 	)

		// }

		if incomingMessage.Type == "wsGetUsers" {
			// saveMessage(
			// 	database,
			// 	incomingMessage.Body,
			// 	incomingMessage.User_id,
			// 	incomingMessage.Target_id,
			// 	incomingMessage.Creation_time,
			// )
			log.Println("Got wsGetUsers from frontend")

			// log.Println(json2.Unmarshal(readUsers(database), nil))
			log.Println(string(readUsers(database)))

		}

		// log.Println(tempMessages)
		// for i := 0; i < len(tempMessages); i++ {
		// 	log.Println(tempMessages[i])
		// }

		// this the action with incoming message, rewrite to func -> db
		log.Println("Printing out received message: ")
		log.Println(string(p))

		// this repeats incoming message back to frontend
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

	// log.Println("Printing out received message: ")
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // avoid CORS error

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Backend: Client Successfully Connected to WebSocket...THE FUCK?")
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
		ID int `json:"id"`
		// Email        string `json:"email,omitempty"`
		// Gender       string `json:"gender,omitempty"`
		// FirstName    string `json:"first_name,omitempty"`
		// LastName     string `json:"last_name,omitempty"`
		Username string `json:"username"`
		// PasswordHash string `json:"-"`
		// CreationTime string `json:"creation_time"`
		LoginTime string `json:"login_Time"`
		// IsAdmin      string `json:"is_admin"`
		// Token        string `json:"-"`
	}
	var data []Wsuser
	var json []byte
	var err error

	// Variables to use for assignment from database
	var userID int
	// var email string
	// var gender string
	// var firstName string
	// var lastName string
	var username string
	// var passwordHash string
	// var createdDate string
	var loginDate string
	// var isAdmin string

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
			ID: userID,
			// Email:        email,
			// Gender:       gender,
			// FirstName:    firstName,
			// LastName:     lastName,
			Username: username,
			// PasswordHash: passwordHash,
			// CreationTime: createdDate,
			LoginTime: loginDate,
			// IsAdmin:      isAdmin,
		}
		rows.Scan(&userID /* &email, &gender, &firstName, &lastName, */, &username /* &passwordHash, &createdDate, */, &loginDate /* , &isAdmin */)

		data = append(data, user)
	}

	if len(data) == 0 {
		logger.WarningLogger.Println("There are 0 users")
	}

	// Write json to return
	json, err = json2.Marshal(data)
	if err != nil {
		logger.ErrorLogger.Println(err)
		// w.WriteHeader(http.StatusBadRequest)
		// return
	}

	return json

}

// func albumsByArtist(artist string) ([]Album, error) {
// }

// var tempMessages = []string{}

// func saveMessage(db *sql.DB, body string, user_id string, target_id string, creation_time string) {

// 	message := `INSERT INTO messages(body, user_id, target_id, creation_time) VALUES (?, ?, ?, ?)`
// 	query, err := db.Prepare(message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = query.Exec(body, user_id, target_id, creation_time)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
