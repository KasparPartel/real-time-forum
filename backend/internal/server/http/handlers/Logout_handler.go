package handlers

import (
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		helper.EnableCors(&w)

		// Extract id from URL
		id := helper.ExtractURLID(r, "logout")

		// Connect to database
		db, err := db2.OpenDB()
		helper.CheckError(err)
		defer db.Close()

		timeNow := time.Now().Format(LongForm)

		_, err = db.Exec("UPDATE user SET token=?, logout_date=? WHERE id=?", "", timeNow, id)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
