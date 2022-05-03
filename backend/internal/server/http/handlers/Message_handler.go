package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
	"strconv"
	"time"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Endpoint hit: api/message")

	// Set correct headers so client can request data
	// Without correct headers there can be CORS errors etc.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// Extract id from URL
	id := helper.ExtractURLID(r, "message")

	// All data from POST response body must be parsed to work with it
	err := r.ParseForm()
	if err != nil {
		return
	}

	// Connect to database
	db, err := db2.Open()
	helper.CheckError(err)
	defer db.Close()

	// Variables to use for assignment from database
	var messageID int
	var body string
	var userID int
	var targetID int
	var createdDate string

	// Time formatting string
	const longForm = "2006-01-02 15:04:05.000 -0700 PDT"

	// Switch over request method - POST, GET, DELETE, UPDATE
	switch r.Method {
	case "POST":
		// If there is id in URI then update a specific post
		// Else create a new post
		if len(id) != 0 {
			logger.InfoLogger.Println("POST: modify a message with form data")

			row := db.QueryRow("SELECT * FROM post WHERE message_id=?", id)

			if err = row.Scan(&messageID, &body, &userID, &targetID, &createdDate); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Message with id %s does not exist", id)
			} else {
				message := model.Message{
					ID:           messageID,
					Body:         r.FormValue("body"),
					UserID:       userID,
					TargetID:     targetID,
					CreationTime: createdDate,
				}

				_, err := db.Exec("UPDATE message SET body=? WHERE post_id=?",
					message.Body, message.ID)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}
			}
		} else {
			logger.InfoLogger.Println("POST: create a message with form data")

			var lastId int

			row := db.QueryRow("SELECT message_id FROM message ORDER BY message_id DESC limit 1")

			if err = row.Scan(&lastId); err == sql.ErrNoRows {
				logger.InfoLogger.Println("No messages found")
			}

			message := model.Message{
				ID: lastId + 1,
				//Title:        r.FormValue("title"),
				Body:     r.FormValue("body"),
				UserID:   1, // ??? why 1
				TargetID: targetID,
				//Filename:     r.FormValue("filename"),
				CreationTime: time.Now().Format(longForm),
				//UpdatedTime:  time.Now().Format(longForm),
			}

			_, err := db.Exec("INSERT INTO message(message_id, body, user_id, target_id, created_date)"+
				"VALUES(?, ?, ?, ?, ?)", message.ID, message.Body, message.UserID, message.TargetID, message.CreationTime)
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
		}

	case "GET":
		logger.InfoLogger.Println("GET")

		var json []byte
		var err error
		var data []model.Message

		// If there is id then return specific message
		// Else return all messages
		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: message with id %s\n", id)

			id, err := strconv.Atoi(id)
			if err != nil {
				return
			}

			row := db.QueryRow("SELECT * FROM message WHERE message_id=?", id)

			if err = row.Scan(&messageID, &body, &userID, &targetID, &createdDate); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Message with id %d does not exist", id)
			} else {
				// post := model.Post{
				// 	ID:           postID,
				// 	Title:        title,
				// 	Body:         body,
				// 	UserID:       userID,
				// 	Filename:     filename,
				// 	CreationTime: createdDate,
				// 	UpdatedTime:  updatedDate,
				// }

				message := model.Message{
					ID: messageID,
					//Title:        r.FormValue("title"),
					Body:     body,
					UserID:   userID,
					TargetID: targetID,
					//Filename:     r.FormValue("filename"),
					CreationTime: createdDate,
					//UpdatedTime:  time.Now().Format(longForm),
				}

				data = append(data, message)
			}

		} else {
			logger.InfoLogger.Println("GET: all messages")

			// Select every row from post table
			rows, err := db.Query("SELECT * FROM message ORDER BY userID") // title --> userID
			helper.CheckError(err)
			defer rows.Close()

			// Loop over every row
			for rows.Next() {
				err := rows.Scan(&messageID, &body, &userID, &targetID, &createdDate)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}

				message := model.Message{
					ID:           messageID,
					Body:         body,
					UserID:       userID,
					TargetID:     targetID,
					CreationTime: createdDate,
				}

				data = append(data, message)
			}

			if len(data) == 0 {
				logger.WarningLogger.Println("There are 0 messages")
			}
		}

		// Write json to Response
		json, err = json2.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write(json)
		if err != nil {
			return
		}

		// case "DELETE":
		// 	// If there is id then delete specific post
		// 	// Else delete all posts
		// 	if len(id) != 0 {
		// 		logger.InfoLogger.Printf("DELETE: post with id %s\n", id)

		// 		_, err := db.Exec("DELETE FROM post WHERE post_id=?", id)
		// 		if err != nil {
		// 			logger.ErrorLogger.Println(err)
		// 		} else {
		// 			logger.InfoLogger.Println("Post deleted")
		// 		}
		// 	} else {
		// 		logger.InfoLogger.Println("DELETE: all posts")

		// 		_, err := db.Exec("DELETE FROM post")
		// 		if err != nil {
		// 			logger.ErrorLogger.Println(err)
		// 		} else {
		// 			logger.InfoLogger.Println("All posts deleted")
		// 		}
		// 	}
	}
}
