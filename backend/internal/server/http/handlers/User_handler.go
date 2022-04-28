package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
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

	// Variables to use for assignment from database
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

	// Switch over request method - POST, GET, DELETE
	switch r.Method {
	case http.MethodPost:
		// Validate form data
		// Email Validation
		if !helper.IsValidEmail(r.FormValue("email")) {
			logger.ErrorLogger.Println("Email is not valid!")
			return
		}
		// Username validation
		if len(r.FormValue("username")) < 5 {
			logger.ErrorLogger.Println("Username must be at least 5 characters long!")
			return
		}
		// Password validation and hashing
		if len(r.FormValue("password")) < 8 {
			logger.ErrorLogger.Println("Password must be at least 5 characters long!")
			return
		}
		passwordHash, err = helper.GeneratePasswordHash(r.FormValue("password"))
		if err != nil {
			logger.ErrorLogger.Println("Cannot hash password!")
			return
		}

		// If there is id in URI then update a specific user
		// Else create a new user - User registration
		if len(id) != 0 {
			logger.InfoLogger.Println("POST: modify a user with form data")

			row := db.QueryRow("SELECT * FROM post WHERE user_id=?", id)

			if err = row.Scan(&userID, &email, &gender, &firstName, &lastName, &username, &passwordHash, &createdDate, &loginDate, isAdmin); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("User with id %d does not exist", id)
			} else {
				user := model.User{
					ID:           userID,
					Email:        r.FormValue("email"),
					Gender:       r.FormValue("gender"),
					FirstName:    r.FormValue("first_name"),
					LastName:     r.FormValue("last_name"),
					Username:     r.FormValue("username"),
					PasswordHash: passwordHash,
					CreationTime: createdDate,
					LoginTime:    loginDate,
					IsAdmin:      isAdmin,
				}

				_, err := db.Exec("UPDATE user SET email=?, gender=?, first_name=?, last_name=?, username=?, password_hash=? WHERE user_id=?",
					user.Email, user.Gender, user.FirstName, user.LastName, user.Username, user.PasswordHash, user.ID)
				if err != nil {
					logger.ErrorLogger.Println(err)
					return
				}
			}
		} else {
			logger.InfoLogger.Println("POST: create a user with form data")

			// Last id from database table
			var lastId int

			row := db.QueryRow("SELECT user_id FROM user ORDER BY user_id DESC limit 1")
			_ = row.Scan(&lastId)

			user := model.User{
				ID:           lastId + 1,
				Email:        r.FormValue("email"),
				Gender:       r.FormValue("gender"),
				FirstName:    r.FormValue("first_name"),
				LastName:     r.FormValue("last_name"),
				Username:     r.FormValue("username"),
				PasswordHash: passwordHash,
				CreationTime: time.Now().Format(longForm),
				LoginTime:    "",
				IsAdmin:      "no",
			}

			_, err := db.Exec("INSERT INTO user(user_id, email, gender, first_name, last_name, username, password_hash, created_date, login_date, administrator)"+
				"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.Email, user.Gender, user.FirstName, user.LastName, user.Username, user.PasswordHash, user.CreationTime, user.LoginTime, user.IsAdmin)
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
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var json []byte
		var err error
		var data []model.User

		// If there is id then return specific post
		// Else return all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: user with id %s\n", id)

			row := db.QueryRow("SELECT * FROM user WHERE user_id=?", id)

			if err = row.Scan(&userID, &email, &gender, &firstName, &lastName, &username, &passwordHash, &createdDate, &loginDate, &isAdmin); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("User with id %d does not exist", id)
			} else {

				user := model.User{
					ID:           userID,
					Email:        email,
					Gender:       gender,
					FirstName:    firstName,
					LastName:     lastName,
					Username:     username,
					PasswordHash: passwordHash,
					CreationTime: createdDate,
					LoginTime:    loginDate,
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

				user := model.User{
					ID:           userID,
					Email:        email,
					Gender:       gender,
					FirstName:    firstName,
					LastName:     lastName,
					Username:     username,
					PasswordHash: passwordHash,
					CreationTime: createdDate,
					LoginTime:    loginDate,
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

		//case http.MethodDelete:
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
