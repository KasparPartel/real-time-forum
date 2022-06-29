package db

import (
	"database/sql"
	"log"
	"os"
	config2 "real-time-forum/pkg/config"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Load config from config pkg
var cfg = &config2.Config

func CheckDB() error {
	file, err := os.Open(cfg.DBfilename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// OpenDB opens database conn and returns pointer to database and error
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DBfilename)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateDB() {
	file, err := os.Create(cfg.DBfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const longForm = "2006-01-02 15:04:05.000 -0700 PDT"

	timeNow := time.Now().Format(longForm)

	stmt := `CREATE TABLE user (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		email text NOT NULL UNIQUE,
		gender text NOT NULL,
		age integer NOT NULL,
		first_name text NOT NULL,
		last_name text NOT NULL,
		username text NOT NULL UNIQUE,
		password_hash text NOT NULL,
		created_date text NOT NULL,
		login_date text,
		logout_date text,
		administrator text NOT NULL,
		token text,
		history text);

		INSERT INTO user (email, gender, age, first_name, last_name, username, password_hash, created_date, login_date, logout_date, administrator, token, history)
		VALUES("test@gmail.com", "male", 25, "Test", "Test", "Test1",
		"$2a$10$zTl.sQ6T9JEREXTR3K7z3u/AP53bO.UxIRapugFNTiObfAxNr.Xy2", ?, "", "", "yes", "", "");

		CREATE TABLE post (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		title text NOT NULL,
		body text NOT NULL,
		user_id integer,
		filename text,
		created_date text NOT NULL,
		updated_date text,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE SET NULL );

		CREATE TABLE messages (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		body text NOT NULL,
		user_id text,
		target_id text,
		creation_time text NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE SET NULL,
		FOREIGN KEY(target_id) REFERENCES user(id) ON DELETE SET NULL );`

	_, err = db.Exec(stmt, timeNow)
	if err != nil {
		log.Fatal(err)
	}
}
