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
)

func DislikeHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	logger.InfoLogger.Println("Endpoint hit: api/like")

	// Connect to database
	db, err := db2.OpenDB()
	helper.CheckError(err)
	defer db.Close()

	// Extract id from URL
	id := helper.ExtractURLID(r, "like")

	var likeID int
	var userID int
	var postID int

	switch r.Method {
	case http.MethodPost:
		var dislike map[string]string

		logger.InfoLogger.Println("POST dislike")

		// Read json body into map
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json2.Unmarshal(b, &dislike)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// User id validation - check if input is number and validate user
		userID, err = strconv.Atoi(dislike["user_id"])
		if err != nil {
			logger.ErrorLogger.Printf("User ID - %s - is not a number!\n", dislike["user_id"])
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
		likeID, err = strconv.Atoi(dislike["post_id"])
		if err != nil {
			logger.ErrorLogger.Printf("Post ID - %s - is not a number!\n", dislike["post_id"])
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
			dislike := model.Dislike{
				UserID: userID,
				PostID: postID,
			}

			_, err = db.Exec("INSERT INTO dislike(user_id, post_id)"+
				"VALUES(?, ?)", dislike.UserID, dislike.PostID)
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
		var data []model.Like

		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: dislikes for post with id %s\n", id)

			rows, err := db.Query("SELECT * FROM like WHERE post_id=?", id)
			helper.CheckError(err)
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&likeID, &userID, &postID)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}

				like := model.Like{
					ID:     likeID,
					UserID: userID,
					PostID: postID,
				}

				data = append(data, like)
			}

			if len(data) == 0 {
				logger.InfoLogger.Println("There are 0 dislikes for this post")
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
