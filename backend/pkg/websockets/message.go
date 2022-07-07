package websockets

import (
	"fmt"
	// "io"
	"database/sql"
	"log"
	"strconv"
	"strings"

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

func WsReadUsers(db *sql.DB) ([]byte, []int) {

	type Wsuser struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		LoginDate  string `json:"login_date"`
		LogoutDate string `json:"logout_date"`
		History    string `json:"history"`
	}
	var data []Wsuser
	var json []byte
	var userArray []int
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
			History:    history,
		}

		data = append(data, user)
		userArray = append(userArray, user.ID)
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

	return json, userArray

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

func WsReadMessages(db *sql.DB, messageUser string, messageTarget string) ([]byte, int) {

	type Wsmessage struct {
		ID            int    `json:"id"`
		Body          string `json:"body"`
		User_id       string `json:"user_id"`
		Target_id     string `json:"target_id"`
		Creation_time string `json:"creation_time"`
	}

	// returnedmessages := []byte(fmt.Sprintf(`{"type":"wsReturnedMessages","sender":"%s","body":`, messageUser))

	var data []Wsmessage
	var json []byte
	var err error

	// Variables to use for assignment from database
	var msgID int
	var msgBody string
	var msgUser string
	var msgTarget string
	var msgCreationTime string

	// logger.InfoLogger.Println("GET: all messages with current user and target")

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

	messagelength := len(data)

	// log.Println("WsReadMessages length:", len(data))

	// returnedmessages = append(returnedmessages, json...)
	// returnedmessages = append(returnedmessages, []byte(`}`)...)

	// return returnedmessages
	return json, messagelength

}

func WsSaveHistory(db *sql.DB, user string, target string) {
	// this function saves a message target into user's chat history array
	// saves into both chat users' history
	// if id already present in history, moves it to start of array

	// 1. query chat user and target history strings from table
	type History struct {
		ID   int
		Data string
	}
	var userInt int
	var targetInt int
	var userHistory History
	var targetHistory History

	if i, err := strconv.Atoi(user); err == nil {
		userInt = i
	}
	if j, err := strconv.Atoi(target); err == nil {
		targetInt = j
	}

	queryString := `SELECT id, history FROM user WHERE id=$1;`

	userrow := db.QueryRow(queryString, userInt)
	switch err := userrow.Scan(&userHistory.ID, &userHistory.Data); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Printf("History of userId %d: %s\n", userHistory.ID, userHistory.Data)
	default:
		panic(err)
	}

	targetrow := db.QueryRow(queryString, targetInt)
	switch err := targetrow.Scan(&targetHistory.ID, &targetHistory.Data); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Printf("History of userId %d: %s\n", targetHistory.ID, targetHistory.Data)
	default:
		panic(err)
	}

	// 2. place user id into start of history
	userHistoryUpdate := convertHistory(userHistory.Data, targetHistory.ID)
	targetHistoryUpdate := convertHistory(targetHistory.Data, userHistory.ID)
	fmt.Printf("New userId %d history: %s\n", userHistory.ID, userHistoryUpdate)
	fmt.Printf("New userId %d history: %s\n", targetHistory.ID, targetHistoryUpdate)

	// 3. save updated strings into user and target history in db
	db.Exec("UPDATE user SET history = ? WHERE id = ?", userHistoryUpdate, userHistory.ID)
	db.Exec("UPDATE user SET history = ? WHERE id = ?", targetHistoryUpdate, targetHistory.ID)
}

func convertHistory(history string, user int) string {
	// this function adds/moves user int to first value in history string

	userStr := strconv.Itoa(user)
	userSplit := strings.Split(history, ",")
	var userRet string
	isfound := false

	for i := 0; i < len(userSplit); i++ {
		if strings.Split(userSplit[i], "-")[0] == userStr {
			isfound = true
		}
	}

	if !isfound {
		userSplit = append([]string{fmt.Sprintf("%s-%s", userStr, "0")}, userSplit...)
	} else {
		for i := 1; i < len(userSplit); i++ {
			if strings.Split(userSplit[i], "-")[0] == userStr {
				userSplit = append([]string{userSplit[i]}, userSplit...)
				if i < len(userSplit)+2 {
					userSplit = append(userSplit[:i+1], userSplit[i+2:]...)
				}
			}
		}
	}

	for i := 0; i < len(userSplit); i++ {
		userRet = userRet + userSplit[i]
		if i < len(userSplit)-1 {
			userRet = userRet + ","
		}
	}

	return userRet
}

func updateHistory(db *sql.DB, history string, user int, target int, length int) /* string */ {
	// this function updates the messages count for target user in history string

	targetStr := strconv.Itoa(target)
	targetSplit := strings.Split(history, ",")
	var historyUpdate string

	for i := 0; i < len(targetSplit); i++ {
		if strings.Split(targetSplit[i], "-")[0] == targetStr {
			targetSplit[i] = fmt.Sprintf("%s-%d", targetStr, length)
		}
		historyUpdate = historyUpdate + targetSplit[i]
		if i < len(targetSplit)-1 {
			historyUpdate = historyUpdate + ","
		}
	}
	db.Exec("UPDATE user SET history = ? WHERE id = ?", historyUpdate, user)
}

func getHistory(db *sql.DB, user int) (int, string) {
	// this func gets user ID and history string from db

	type History struct {
		ID   int
		Data string
	}
	var history History

	queryString := `SELECT id, history FROM user WHERE id=$1;`

	userrow := db.QueryRow(queryString, user)
	switch err := userrow.Scan(&history.ID, &history.Data); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Printf("History of userId %d: %s\n", history.ID, history.Data)
	default:
		panic(err)
	}
	return history.ID, history.Data
}

func compareHistory(history string, user int, length int) bool {
	// this function compares target user messages count with given length count in history string

	userStr := strconv.Itoa(user)
	userSplit := strings.Split(history, ",")
	var messages int

	for i := 0; i < len(userSplit); i++ {
		if strings.Split(userSplit[i], "-")[0] == userStr {
			if i, err := strconv.Atoi(strings.Split(userSplit[i], "-")[1]); err == nil {
				messages = i
			}
			if length > messages {
				return true
			}
		}
	}
	return false
}
