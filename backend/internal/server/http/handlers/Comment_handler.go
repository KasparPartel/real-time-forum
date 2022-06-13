package handlers

import (
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	logger.InfoLogger.Println("Endpoint hit: api/user")

	// Connect to database
	db, err := db2.OpenDB()
	helper.CheckError(err)
	defer db.Close()

	// Extract id from URL
	id := helper.ExtractURLID(r, "user")

	switch r.Method {
	case http.MethodPost:
		var comment map[string]string

	case http.MethodGet:

	}
}
