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

func UserHandler(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Endpoint hit: api/user")

	// Extract id from URL
	id := helper.ExtractURLID(r, "user")

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

		var json []byte
		var err error
		var data []model.User

		var userID int
		var email string
		var gender string
		var firstName string
		var lastName string
		var username string
		var passwordHash string
		var createdDate string
		var loginDate string
		var isAdmin string

		// If there is id then return specific post
		// Else return all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: user with id %s\n", id)

			id, err := strconv.Atoi(id)
			if err != nil {
				return
			}

			row := db.QueryRow("SELECT * FROM user WHERE user_id=?", id)

			if err = row.Scan(&userID, &email, &gender, &firstName, &lastName, &username, &passwordHash, &createdDate, &loginDate, &isAdmin); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("User with id %d does not exist", id)
			} else {
				createdTime, _ := time.Parse(time.RFC3339, createdDate)
				loginTime, _ := time.Parse(time.RFC3339, loginDate)

				user := model.User{
					ID:           userID,
					Email:        email,
					Gender:       gender,
					FirstName:    firstName,
					LastName:     lastName,
					Username:     username,
					PasswordHash: passwordHash,
					CreationTime: createdTime,
					LoginTime:    loginTime,
					IsAdmin:      isAdmin,
				}

				data = append(data, user)
			}
		} else {
			logger.InfoLogger.Println("GET: all users")

			// Select every row from post table
			rows, err := db.Query("SELECT * FROM user ORDER BY first_name")
			helper.CheckError(err)
			defer rows.Close()

			// Loop over every row
			for rows.Next() {
				rows.Scan(&userID, &email, &gender, &firstName, &lastName, &username, &passwordHash, &createdDate, &loginDate, &isAdmin)

				createdTime, _ := time.Parse(time.RFC3339, createdDate)
				loginTime, _ := time.Parse(time.RFC3339, loginDate)

				user := model.User{
					ID:           userID,
					Email:        email,
					Gender:       gender,
					FirstName:    firstName,
					LastName:     lastName,
					Username:     username,
					PasswordHash: passwordHash,
					CreationTime: createdTime,
					LoginTime:    loginTime,
					IsAdmin:      isAdmin,
				}

				data = append(data, user)
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
