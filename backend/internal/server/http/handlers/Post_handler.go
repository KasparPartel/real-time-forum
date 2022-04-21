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

	// Switch over request method - POST, GET, DELETE, UPDATE
	switch r.Method {
	//case "POST":
	//	// If there is id in URI then update a specific post
	//	// Else create a new post
	//	if len(id) != 0 {
	//		id, err := strconv.Atoi(id)
	//		if err != nil {
	//			return
	//		}
	//
	//		logger.InfoLogger.Println("POST: modify a post with form data")
	//
	//		var sliceItemIndex int
	//		var postExists bool
	//
	//		for i, v := range data {
	//			if v.ID == id {
	//				sliceItemIndex = i
	//				postExists = true
	//				break
	//			}
	//		}
	//
	//		if postExists {
	//			data[sliceItemIndex].Title = r.FormValue("title")
	//			data[sliceItemIndex].Body = r.FormValue("body")
	//			data[sliceItemIndex].Filename = r.FormValue("filename")
	//		} else {
	//			logger.ErrorLogger.Printf("Post with id %d does not exist", id)
	//		}
	//	} else {
	//		logger.InfoLogger.Println("POST: create a post with form data")
	//
	//		// Convert data into right format
	//		formID, err := strconv.Atoi(r.FormValue("formID"))
	//		if err != nil {
	//			fmt.Println(formID)
	//			logger.ErrorLogger.Println("formID is not present or wrong format")
	//			return
	//		}
	//
	//		post := model.Post{
	//			ID:           formID,
	//			Title:        r.FormValue("title"),
	//			Body:         r.FormValue("body"),
	//			UserID:       1,
	//			Filename:     r.FormValue("filename"),
	//			CreationTime: time.Now(),
	//			UpdatedTime:  time.Now(),
	//		}
	//
	//		// Append created post into slice
	//		data = append(data, post)
	//	}

	case "GET":
		// Set correct headers so client can request data
		// Without correct headers there can be CORS errors etc.
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		var json []byte
		var err error
		var data []model.Post

		var postID int
		var title string
		var body string
		var userID int
		var filename string
		var createdDate string
		var updatedDate string

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
				createdTime, _ := time.Parse(time.RFC3339, createdDate)
				updatedTime, _ := time.Parse(time.RFC3339, updatedDate)

				post := model.Post{
					ID:           postID,
					Title:        title,
					Body:         body,
					UserID:       userID,
					Filename:     filename,
					CreationTime: createdTime,
					UpdatedTime:  updatedTime,
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
				rows.Scan(&postID, &title, &body, &userID, &filename, &createdDate, &updatedDate)

				createdTime, _ := time.Parse(time.RFC3339, createdDate)
				updatedTime, _ := time.Parse(time.RFC3339, updatedDate)

				post := model.Post{
					ID:           postID,
					Title:        title,
					Body:         body,
					UserID:       userID,
					Filename:     filename,
					CreationTime: createdTime,
					UpdatedTime:  updatedTime,
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

		//case "DELETE":
		//	// If there is id then delete specific post
		//	// Else delete all posts
		//	if len(id) != 0 {
		//		logger.InfoLogger.Printf("DELETE: post with id %s\n", id)
		//
		//		id, err := strconv.Atoi(id)
		//		if err != nil {
		//			return
		//		}
		//
		//		var sliceItemIndex int
		//		var indexSet bool
		//
		//		for i, v := range data {
		//			if v.ID == id {
		//				sliceItemIndex = i
		//				indexSet = true
		//				break
		//			}
		//		}
		//
		//		if indexSet {
		//			data = append(data[:sliceItemIndex], data[sliceItemIndex+1:]...)
		//			logger.ErrorLogger.Printf("Post with id %d deleted\n", id)
		//		} else {
		//			logger.ErrorLogger.Printf("Cannot find post with id %d\n", id)
		//		}
		//
		//	} else {
		//		logger.InfoLogger.Println("DELETE: all posts")
		//
		//		data = nil
		//
		//		logger.InfoLogger.Println("All posts deleted")
		//	}
	}
}
