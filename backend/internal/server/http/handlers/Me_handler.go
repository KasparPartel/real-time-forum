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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	var tokenCookie *http.Cookie
	var err error

	if tokenCookie, err = r.Cookie("session_token"); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logger.InfoLogger.Println("Token not found")
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
	var token string

	if r.Method == http.MethodGet {
		row := db.QueryRow("SELECT * FROM user WHERE token=?", tokenCookie.Value)

		if err = row.Scan(&userID, &email, &gender, &firstName, &lastName, &username, &passwordHash, &createdDate, &loginDate, &isAdmin, &token); err == sql.ErrNoRows {
			logger.ErrorLogger.Printf("User with token %d does not exist", tokenCookie.Value)
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
				Token:        token,
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
