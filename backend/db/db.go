package db

import (
	"database/sql"
	"real-time-forum/pkg/helper"
)

func Open() {
	db, err := sql.Open("sqlite3", "./forum.db")
	helper.CheckError(err)

	defer db.Close()
}
