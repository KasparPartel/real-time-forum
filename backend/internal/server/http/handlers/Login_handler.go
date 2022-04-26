package handlers

import (
	"database/sql"
	uuid2 "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/model"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_token")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		var userID int
		var passwordHash string
		var password = r.FormValue("password")

		// Connect to database
		db, err := db2.Open()
		helper.CheckError(err)
		defer db.Close()

		row := db.QueryRow("SELECT user_id, password_hash FROM user WHERE username=? OR email=?",
			r.FormValue("username"), r.FormValue("email"))

		if err = row.Scan(&userID, &passwordHash); err == sql.ErrNoRows {
			http.Error(w, "User with this username/email does not exist", http.StatusForbidden)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// Create session
		sessionToken := uuid2.NewV4().String()
		expiresAt := time.Now().Add(600 * time.Second)

		model.Sessions[sessionToken] = model.Session{
			UserID:     userID,
			ExpiryTime: expiresAt,
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "login_token",
			Value:  sessionToken,
			MaxAge: 600,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
