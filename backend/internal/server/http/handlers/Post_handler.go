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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Endpoint hit: api/post")

	// Set correct headers so client can request data
	// Without correct headers there can be CORS errors etc.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// Extract id from URL
	id := helper.ExtractURLID(r, "post")

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
	var postID int
	var title string
	var body string
	var userID int
	var filename string
	var createdDate string
	var updatedDate string

	// Time formatting string
	const longForm = "2006-01-02 15:04:05.000 -0700 PDT"

	// Switch over request method - POST, GET, DELETE, UPDATE
	switch r.Method {
	case "POST":
		// If there is id in URI then update a specific post
		// Else create a new post
		if len(id) != 0 {
			logger.InfoLogger.Println("POST: modify a post with form data")

			row := db.QueryRow("SELECT * FROM post WHERE post_id=?", id)

			if err = row.Scan(&postID, &title, &body, &userID, &filename, &createdDate, &updatedDate); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Post with id %d does not exist", id)
			} else {
				post := model.Post{
					ID:           postID,
					Title:        r.FormValue("title"),
					Body:         r.FormValue("title"),
					UserID:       userID,
					Filename:     r.FormValue("filename"),
					CreationTime: createdDate,
					UpdatedTime:  time.Now().Format(longForm),
				}

				_, err := db.Exec("UPDATE post SET title=?, body=?, filename=?, updated_date=? WHERE post_id=?",
					post.Title, post.Body, post.Filename, post.UpdatedTime, post.ID)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}
			}
		} else {
			logger.InfoLogger.Println("POST: create a post with form data")

			// lastId
			var lastId int

			row := db.QueryRow("SELECT post_id FROM post ORDER BY post_id DESC limit 1")

			if err = row.Scan(&lastId); err == sql.ErrNoRows {
				logger.InfoLogger.Println("No posts found")
			}

			post := model.Post{
				ID:           lastId + 1,
				Title:        r.FormValue("title"),
				Body:         r.FormValue("body"),
				UserID:       1,
				Filename:     r.FormValue("filename"),
				CreationTime: time.Now().Format(longForm),
				UpdatedTime:  time.Now().Format(longForm),
			}

			_, err := db.Exec("INSERT INTO post(post_id, title, body, user_id, filename, created_date, updated_date)"+
				"VALUES(?, ?, ?, ?, ?, ?, ?)", post.ID, post.Title, post.Body, post.UserID, post.Filename, post.CreationTime, post.UpdatedTime)
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
		}

	case "GET":
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

			row := db.QueryRow("SELECT * FROM post WHERE post_id=?", id)

			if err = row.Scan(&postID, &title, &body, &userID, &filename, &createdDate, &updatedDate); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Post with id %d does not exist", id)
			} else {
				post := model.Post{
					ID:           postID,
					Title:        title,
					Body:         body,
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
				err := rows.Scan(&postID, &title, &body, &userID, &filename, &createdDate, &updatedDate)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}

				post := model.Post{
					ID:           postID,
					Title:        title,
					Body:         body,
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
			fmt.Println(err)
		}
		_, err = w.Write(json)
		if err != nil {
			return
		}

	case "DELETE":
		// If there is id then delete specific post
		// Else delete all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("DELETE: post with id %s\n", id)

			_, err := db.Exec("DELETE FROM post WHERE post_id=?", id)
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
	}
}
