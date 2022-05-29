package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"io/ioutil"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
	"strconv"
	"time"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)

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
	var age int
	// var age string // ET: trying to fix User_handler.go

	// Switch over request method - POST, GET, DELETE
	switch r.Method {
	case http.MethodPost:
		var register map[string]string

		// Read body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// All data from POST response body must be parsed to work with it
		err = r.ParseForm()
		if err != nil {
			return
		}

		err = json2.Unmarshal(b, &register)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Validate form data
		// Username validation
		if len(register["username"]) < 5 {
			logger.ErrorLogger.Println("Username must be at least 5 characters long!")
			http.Error(w, "Username must be at least 5 characters!", http.StatusBadRequest)
			return
		}
		// Password validation and hashing
		if len(register["password"]) < 8 {
			logger.ErrorLogger.Println("Password must be at least 5 characters long!")
			http.Error(w, "Password must be at least 5 characters!", http.StatusBadRequest)
			return
		}
		age, err = strconv.Atoi(register["age"])
		if err != nil {
			logger.ErrorLogger.Println("Age is not a number!")
			http.Error(w, "Age is not a number!", http.StatusBadRequest)
			return
		}
		passwordHash, err = helper.GeneratePasswordHash(register["password"])
		if err != nil {
			logger.ErrorLogger.Println("Cannot hash password!")
			return
		}

		// If there is id in URI then update a specific user
		// Else create a new user - User registration
		if len(id) != 0 {
			logger.InfoLogger.Println("POST: modify a user with form data")

			row := db.QueryRow("SELECT * FROM post WHERE user_id=?", id)

			if err = row.Scan(&userID, &email, &gender, &firstName, &lastName, &age, &username, &passwordHash, &createdDate, &loginDate, isAdmin); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("User with id %d does not exist", id)
			} else {
				user := model.User{
					ID:        userID,
					Email:     register["email"],
					Gender:    register["gender"],
					FirstName: register["first_name"],
					LastName:  register["last_name"],
					Age:       age,
					// Age:          register["age"],
					Username:     register["username"],
					PasswordHash: passwordHash,
					CreationTime: createdDate,
					LoginTime:    loginDate,
					IsAdmin:      isAdmin,
				}

				_, err := db.Exec("UPDATE user SET email=?, gender=?, first_name=?, age=?, last_name=?, username=?, password_hash=? WHERE user_id=?",
					user.Email, user.Gender, user.FirstName, user.LastName, user.Age, user.Username, user.PasswordHash, user.ID)
				if err != nil {
					logger.ErrorLogger.Println(err)
					return
				}
			}
		} else {
			logger.InfoLogger.Println("POST: create a user with form data")

			var myId string
			if err = db.QueryRow("SELECT user_id FROM user WHERE username=?", register["username"]).Scan(&myId); err == nil {
				logger.ErrorLogger.Println(err)
				logger.InfoLogger.Println("User with this username already exists!")
				http.Error(w, "User with this username already exists!", http.StatusBadRequest)
				return
			}
			if err = db.QueryRow("SELECT user_id FROM user WHERE email=?", register["email"]).Scan(&myId); err == nil {
				logger.InfoLogger.Println("User with this email already exists!")
				http.Error(w, "User with this email already exists!", http.StatusBadRequest)
				return
			}

			// Last id from database table
			var lastId int

			row := db.QueryRow("SELECT user_id FROM user ORDER BY user_id DESC limit 1")
			_ = row.Scan(&lastId)

			user := model.User{
				ID:        lastId + 1,
				Email:     register["email"],
				Gender:    register["gender"],
				FirstName: register["first_name"],
				LastName:  register["last_name"],
				Age:       age,
				// Age:          register["age"],
				Username:     register["username"],
				PasswordHash: passwordHash,
				CreationTime: time.Now().Format(longForm),
				LoginTime:    "",
				IsAdmin:      "no",
			}

			_, err := db.Exec("INSERT INTO user(user_id, email, gender, first_name, last_name, age, username, password_hash, created_date, login_date, administrator)"+
				"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.Email, user.Gender, user.FirstName, user.LastName, user.Age, user.Username, user.PasswordHash, user.CreationTime, user.LoginTime, user.IsAdmin)
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

			if err = row.Scan(&userID, &email, &gender, &firstName, &lastName, &age, &username, &passwordHash, &createdDate, &loginDate, &isAdmin); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("User with id %d does not exist", id)
			} else {

				user := model.User{
					ID:           userID,
					Email:        email,
					Gender:       gender,
					FirstName:    firstName,
					LastName:     lastName,
					Age:          age,
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
				rows.Scan(&userID, &email, &gender, &firstName, &lastName, &age, &username, &passwordHash, &createdDate, &loginDate, &isAdmin)

				user := model.User{
					ID:           userID,
					Email:        email,
					Gender:       gender,
					FirstName:    firstName,
					LastName:     lastName,
					Age:          age,
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
