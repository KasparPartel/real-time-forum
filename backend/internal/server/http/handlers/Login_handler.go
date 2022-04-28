package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	uuid2 "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		type Login struct {
			Username string
			Email    string
			Password string
		}
		var login Login

		var userID int
		var passwordHash string
		//var password = r.FormValue("password")

		// All data from POST response body must be parsed to work with it
		err := r.ParseForm()
		if err != nil {
			return
		}

		_ = json.Unmarshal([]byte(r.Form), &login)
		fmt.Println(r.Form)

		// Connect to database
		db, err := db2.Open()
		helper.CheckError(err)
		defer db.Close()

		row := db.QueryRow("SELECT user_id, password_hash FROM user WHERE username=? OR email=?",
			r.FormValue("username"), r.FormValue("email"))

		if err = row.Scan(&userID, &passwordHash); err == sql.ErrNoRows {
			http.Error(w, "User with this username/email does not exist", http.StatusForbidden)
			logger.WarningLogger.Println("User with this username/email does not exist")
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			logger.WarningLogger.Println("Username and/or password do not match")
			return
		}

		// Create session
		sessionToken := uuid2.NewV4().String()
		timeNow := time.Now().Format(longForm)

		_, err = db.Exec("UPDATE user SET login_time=?, token=? WHERE user_id=?",
			timeNow, sessionToken, userID)

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			MaxAge:   600,
			HttpOnly: true,
		})

		w.WriteHeader(http.StatusOK)
	}
}
