package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
	"strconv"
	"time"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	logger.InfoLogger.Println("Endpoint hit: api/comment")

	logger.InfoLogger.Println(r.URL)

	// Connect to database
	db, err := db2.OpenDB()
	helper.CheckError(err)
	defer db.Close()

	// Extract id from URL
	id := helper.ExtractURLID(r, "comments")

	var commentID int
	var userID int
	var postID int
	var body string
	var createdDate string

	switch r.Method {
	case http.MethodPost:
		var comment map[string]string

		logger.InfoLogger.Println("POST comment")

		// Read json body into map
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json2.Unmarshal(b, &comment)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// User id validation - check if input is number and validate user
		userID, err = strconv.Atoi(comment["user_id"])
		if err != nil {
			logger.ErrorLogger.Printf("User ID - %s - is not a number!\n", comment["user_id"])
			http.Error(w, "User ID is not a number!", http.StatusBadRequest)
			return
		}
		row := db.QueryRow("SELECT id FROM user WHERE id=?", userID)
		if err = row.Scan(&userID); err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("User with id %d does not exist\n", userID)
			http.Error(w, fmt.Sprintf("User with id %d does not exist\n", userID), http.StatusBadRequest)
			return
		}
		// Post id validation - check if input is number and validate post
		postID, err = strconv.Atoi(comment["post_id"])
		if err != nil {
			logger.ErrorLogger.Printf("Post ID - %s - is not a number!\n", comment["post_id"])
			http.Error(w, "Post ID is not a number!", http.StatusBadRequest)
			return
		}
		row = db.QueryRow("SELECT id FROM post WHERE id=?", postID)
		if err = row.Scan(&postID); err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("Post with id %d does not exist\n", postID)
			http.Error(w, fmt.Sprintf("Post with id %d does not exist\n", postID), http.StatusBadRequest)
			return
		}

		if len(id) != 0 {

		} else {
			comment := model.Comment{
				UserID:       userID,
				PostID:       postID,
				Body:         comment["body"],
				CreationTime: time.Now().Format(LongForm),
			}

			_, err = db.Exec("INSERT INTO comment(user_id, post_id, body, created_date)"+
				"VALUES(?, ?, ?, ?)", comment.UserID, comment.PostID, comment.Body, comment.CreationTime)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
		}

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		var json []byte
		var data []model.Comment

		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: comments for post with id %s\n", id)

			rows, err := db.Query("SELECT * FROM comment WHERE post_id=?", id)
			helper.CheckError(err)
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&commentID, &userID, &postID, &body, &createdDate)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}

				comment := model.Comment{
					ID:           commentID,
					UserID:       userID,
					PostID:       postID,
					Body:         body,
					CreationTime: createdDate,
				}

				data = append(data, comment)
			}

			if len(data) == 0 {
				logger.InfoLogger.Println("There are 0 comments for this post")
			}

			// Write json to Response
			json, err = json2.Marshal(data)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			_, err = w.Write(json)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		}
	}
}
