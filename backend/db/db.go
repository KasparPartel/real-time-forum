package db

import (
	"database/sql"
	config2 "real-time-forum/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

// Open opens database conn and returns pointer to database and error
func Open() (*sql.DB, error) {
	// Load config from config pkg
	cfg := &config2.Config

	db, err := sql.Open("sqlite3", cfg.DBfilename)
	if err != nil {
		return nil, err
	}

	return db, nil
}
