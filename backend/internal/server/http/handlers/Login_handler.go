package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	uuid2 "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"time"
)

type Login struct {
	Username string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if _, err := r.Cookie("session_token"); err == nil {
		w.WriteHeader(http.StatusSeeOther)
		logger.InfoLogger.Println("User already logged in!")
		return
	}

	if r.Method == http.MethodPost {
		// Read body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var login Login
		var userID int
		var passwordHash string

		// All data from POST response body must be parsed to work with it
		err = r.ParseForm()
		if err != nil {
			return
		}

		err = json.Unmarshal(b, &login)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Connect to database
		db, err := db2.OpenDB()
		helper.CheckError(err)
		defer db.Close()

		row := db.QueryRow("SELECT id, password_hash FROM user WHERE username=? OR email=?",
			login.Username, login.Username)

		if err = row.Scan(&userID, &passwordHash); err == sql.ErrNoRows {
			http.Error(w, "User with this username/email does not exist", http.StatusForbidden)
			logger.InfoLogger.Println("User with this username/email does not exist")
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(login.Password)); err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			logger.InfoLogger.Println("Username and/or password do not match")
			return
		}

		// Create session
		sessionToken := uuid2.NewV4().String()
		timeNow := time.Now().Format(LongForm)

		_, err = db.Exec("UPDATE user SET login_date=?, token=? WHERE id=?",
			timeNow, sessionToken, userID)
		if err != nil {
			http.Error(w, "Error writing to database", http.StatusInternalServerError)
			logger.ErrorLogger.Println(err)
			return
		}

		token := map[string]string{"token": sessionToken}

		jsonToken, err := json.Marshal(token)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Write(jsonToken)

	}
}
