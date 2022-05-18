package handlers

import (
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)

	// Extract id from URL
	id := helper.ExtractURLID(r, "logout")

	// Connect to database
	db, err := db2.Open()
	helper.CheckError(err)
	defer db.Close()

	if r.Method == http.MethodPost {
		_, err := db.Exec("UPDATE user SET token=? WHERE user_id=?", "", id)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
