package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
)

func MeHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var tokenCookie *http.Cookie
	var err error

	if tokenCookie, err = r.Cookie("session_token"); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logger.InfoLogger.Println("Token not found")
		return
	}

	// Connect to database
	db, err := db2.OpenDB()
	helper.CheckError(err)
	defer db.Close()

	// Variables to use for assignment from database
	var userID int
	var email string
	var gender string
	var age int
	var firstName string
	var lastName string
	var username string
	var passwordHash string
	var createdDate string
	var loginDate string
	var logoutDate string
	var isAdmin string
	var token string
	var history string

	if r.Method == http.MethodGet {
		row := db.QueryRow("SELECT * FROM user WHERE token=?", tokenCookie.Value)

		if err = row.Scan(&userID, &email, &gender, &age, &firstName, &lastName, &username, &passwordHash, &createdDate, &loginDate, &logoutDate, &isAdmin, &token, &history); err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("User with token %s does not exist", tokenCookie.Value)
		} else {
			user := model.User{
				ID:           userID,
				Email:        email,
				Gender:       gender,
				Age:          age,
				FirstName:    firstName,
				LastName:     lastName,
				Username:     username,
				PasswordHash: passwordHash,
				CreationTime: createdDate,
				LoginTime:    loginDate,
				LogoutTime:   logoutDate,
				IsAdmin:      isAdmin,
				Token:        token,
				History:      history,
			}

			json, err := json2.Marshal(user)
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
