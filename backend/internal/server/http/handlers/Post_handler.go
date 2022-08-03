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

// LongForm - Time formatting string
const LongForm = "2006-01-02 15:04:05.000 -0700 PDT"

func PostHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	logger.InfoLogger.Println("Endpoint hit: api/post")

	// Extract id from URL
	id := helper.ExtractURLID(r, "post")

	// All data from POST response body must be parsed to work with it
	err := r.ParseForm()
	if err != nil {
		return
	}

	// Connect to database
	db, err := db2.OpenDB()
	helper.CheckError(err)
	defer db.Close()

	// Variables to use for assignment from database
	var postID int
	var title string
	var body string
	var categoryID int
	var userID int
	var filename string
	var createdDate string
	var updatedDate string

	// Switch over request method - POST, GET, DELETE
	switch r.Method {
	case http.MethodPost:
		var post map[string]string

		// Read json body into map
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json2.Unmarshal(b, &post)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		/* Validate form data */
		// User id validation - check if input is number and validate user
		userID, err = strconv.Atoi(post["user_id"])
		if err != nil {
			logger.ErrorLogger.Printf("User ID - %s - is not a number!\n", post["user_id"])
			http.Error(w, "User ID is not a number!", http.StatusBadRequest)
			return
		}
		row := db.QueryRow("SELECT id FROM user WHERE id=?", userID)
		if err = row.Scan(&userID); err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("User with id %d does not exist\n", userID)
			http.Error(w, fmt.Sprintf("User with id %d does not exist\n", userID), http.StatusBadRequest)
			return
		}

		// Category id validation - check if input is number and validate category
		categoryID, err = strconv.Atoi(post["category_id"])
		if err != nil {
			logger.ErrorLogger.Printf("Category ID - %s - is not a number!\n", post["category_id"])
			http.Error(w, "Category ID is not a number!", http.StatusBadRequest)
			return
		}
		row = db.QueryRow("SELECT id FROM category WHERE id=?", categoryID)
		if err = row.Scan(&categoryID); err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("Category with id %d does not exist\n", categoryID)
			http.Error(w, fmt.Sprintf("Category with id %d does not exist\n", categoryID), http.StatusBadRequest)
			return
		}

		// If there is id in URI then update a specific post
		// Else create a new post
		if len(id) != 0 {
			logger.InfoLogger.Println("POST: modify a post with form data")

			row := db.QueryRow("SELECT * FROM post WHERE id=?", id)

			if err = row.Scan(&postID, &title, &body, &userID, &filename, &createdDate, &updatedDate); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Post with id %s does not exist", id)
			} else {
				post := model.Post{
					ID:           postID,
					Title:        post["title"],
					Body:         post["body"],
					CategoryID:   categoryID,
					UserID:       userID,
					Filename:     post["filename"],
					CreationTime: createdDate,
					UpdatedTime:  time.Now().Format(LongForm),
				}

				_, err := db.Exec("UPDATE post SET title=?, body=?, filename=?, updated_date=? WHERE id=?",
					post.Title, post.Body, post.Filename, post.UpdatedTime, post.ID)
				if err != nil {
					logger.ErrorLogger.Println(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
		} else {
			logger.InfoLogger.Println("POST: create a post with form data")

			post := model.Post{
				Title:        post["title"],
				Body:         post["body"],
				CategoryID:   categoryID,
				UserID:       userID,
				Filename:     post["filename"],
				CreationTime: time.Now().Format(LongForm),
				UpdatedTime:  time.Now().Format(LongForm),
			}

			_, err := db.Exec("INSERT INTO post(title, body, category_id, user_id, filename, created_date, updated_date)"+
				"VALUES(?, ?, ?, ?, ?, ?, ?)", post.Title, post.Body, post.CategoryID,
				post.UserID, post.Filename, post.CreationTime, post.UpdatedTime)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			row := db.QueryRow("SELECT id FROM post WHERE user_id=? ORDER BY id DESC LIMIT 1", userID)
			if err = row.Scan(&id); err == sql.ErrNoRows {
				logger.ErrorLogger.Println("Post not found")
				http.Error(w, "Post not found", http.StatusInternalServerError)
				return
			}

			// Send back post id for redirect
			json, err := json2.Marshal(id)
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
		}

		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:
		// Set correct headers so client can request data
		// Without correct headers there can be CORS errors etc.
		w.Header().Set("Content-Type", "application/json")

		var json []byte
		var err error
		var data []model.Post

		// If there is id then return specific post
		// Else return all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: post with id %s\n", id)

			id, err := strconv.Atoi(id)
			if err != nil {
				return
			}

			row := db.QueryRow("SELECT * FROM post WHERE id=?", id)

			if err = row.Scan(&postID, &title, &body, &categoryID, &userID, &filename, &createdDate, &updatedDate); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Post with id %d does not exist", id)
			} else {
				post := model.Post{
					ID:           postID,
					Title:        title,
					Body:         body,
					CategoryID:   categoryID,
					UserID:       userID,
					Filename:     filename,
					CreationTime: createdDate,
					UpdatedTime:  updatedDate,
				}

				data = append(data, post)
			}

		} else {
			logger.InfoLogger.Println("GET: all posts")

			// Select every row from post table
			rows, err := db.Query("SELECT * FROM post ORDER BY title")
			helper.CheckError(err)
			defer rows.Close()

			// Loop over every row
			for rows.Next() {
				err := rows.Scan(&postID, &title, &body, &categoryID, &userID, &filename, &createdDate, &updatedDate)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}

				post := model.Post{
					ID:           postID,
					Title:        title,
					Body:         body,
					CategoryID:   categoryID,
					UserID:       userID,
					Filename:     filename,
					CreationTime: createdDate,
					UpdatedTime:  updatedDate,
				}

				data = append(data, post)
			}

			if len(data) == 0 {
				logger.WarningLogger.Println("There are 0 posts")
			}
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

	case http.MethodDelete:
		// If there is id then delete specific post
		// Else delete all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("DELETE: post with id %s\n", id)

			_, err := db.Exec("DELETE FROM post WHERE id=?", id)
			if err != nil {
				logger.ErrorLogger.Println(err)
			} else {
				logger.InfoLogger.Println("Post deleted")
			}
		} else {
			logger.InfoLogger.Println("DELETE: all posts")

			_, err := db.Exec("DELETE FROM post")
			if err != nil {
				logger.ErrorLogger.Println(err)
			} else {
				logger.InfoLogger.Println("All posts deleted")
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}
