package handlers

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	db2 "real-time-forum/db"
	"real-time-forum/pkg/helper"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
	"strconv"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(&w)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	logger.InfoLogger.Println("Endpoint hit: api/category")

	// Connect to database
	db, err := db2.OpenDB()
	helper.CheckError(err)
	defer db.Close()

	// Extract id from URL
	id := helper.ExtractURLID(r, "categories")

	var categoryID int
	var title string

	switch r.Method {
	case http.MethodPost:
		var category map[string]string

		logger.InfoLogger.Println("POST: category")

		// Read json body into map
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json2.Unmarshal(b, &category)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Category id validation - check if input is number
		categoryID, err = strconv.Atoi(category["category_id"])
		if err != nil {
			logger.ErrorLogger.Printf("Category ID - %s - is not a number!\n", category["category_id"])
			http.Error(w, "Category ID is not a number!", http.StatusBadRequest)
			return
		}

		if len(id) != 0 {

		} else {
			category := model.Category{
				Title: category["title"],
			}

			_, err = db.Exec("INSERT INTO category(title)"+
				"VALUES(?)", category.Title)
			if err != nil {
				logger.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
		}

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		var json []byte
		var data []model.Category

		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: category with id %s\n", id)

			row := db.QueryRow("SELECT id, title FROM category WHERE id=?", id)
			if err = row.Scan(&categoryID, &title); err == sql.ErrNoRows {
				logger.ErrorLogger.Printf("Category %s does not exist", title)
				http.Error(w, fmt.Sprintf("Category %s does not exist", title), http.StatusBadRequest)
			}

			category := model.Category{
				ID:    categoryID,
				Title: title,
			}

			data = append(data, category)

		} else {
			rows, err := db.Query("SELECT * FROM category ORDER BY id")
			helper.CheckError(err)
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&categoryID, &title)
				if err != nil {
					logger.ErrorLogger.Println(err)
				}

				category := model.Category{
					ID:    categoryID,
					Title: title,
				}

				data = append(data, category)
			}

		}

		if len(data) == 0 {
			logger.InfoLogger.Println("There is no category for this post")
		}
		// Write json to Response
		json, err = json2.Marshal(data)
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
