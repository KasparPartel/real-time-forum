package handlers

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/db"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/model"
	"strconv"
	"strings"
	"time"
)

var data []model.Post

func PostHandler(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Endpoint hit: api/post")

	// Extract id from URL
	id := strings.TrimPrefix(r.URL.Path, "/api/post/")

	// All data from POST response body must be parsed to work with it
	err := r.ParseForm()
	if err != nil {
		return
	}

	// Switch over request method - POST, GET, DELETE, UPDATE
	switch r.Method {
	case "POST":
		// If there is id in URI then update a specific post
		// Else create a new post
		if len(id) != 0 {
			id, err := strconv.Atoi(id)
			if err != nil {
				return
			}

			logger.InfoLogger.Println("POST: modify a post with form data")

			var sliceItemIndex int
			var postExists bool

			for i, v := range data {
				if v.ID == id {
					sliceItemIndex = i
					postExists = true
					break
				}
			}

			if postExists {
				data[sliceItemIndex].Title = r.FormValue("title")
				data[sliceItemIndex].Body = r.FormValue("body")
				data[sliceItemIndex].Filename = r.FormValue("filename")
			} else {
				logger.ErrorLogger.Printf("Post with id %d does not exist", id)
			}
		} else {
			logger.InfoLogger.Println("POST: create a post with form data")

			// Convert data into right format
			formID, err := strconv.Atoi(r.FormValue("formID"))
			if err != nil {
				fmt.Println(formID)
				logger.ErrorLogger.Println("formID is not present or wrong format")
				return
			}
			formLikeAmount, err := strconv.Atoi(r.FormValue("formLikeAmount"))
			if err != nil {
				logger.ErrorLogger.Println("formLikeAmount is not present or wrong format")
				return
			}
			formDislikeAmount, err := strconv.Atoi(r.FormValue("formDislikeAmount"))
			if err != nil {
				logger.ErrorLogger.Println("formDislikeAmount is not present or wrong format")
				return
			}
			formCommentAmount, err := strconv.Atoi(r.FormValue("formCommentAmount"))
			if err != nil {
				logger.ErrorLogger.Println("formCommentAmount is not present or wrong format")
				return
			}

			post := model.Post{
				ID:            formID,
				Title:         r.FormValue("title"),
				Body:          r.FormValue("body"),
				Author:        db.DummyUsers[0],
				Filename:      r.FormValue("filename"),
				LikeAmount:    formLikeAmount,
				DislikeAmount: formDislikeAmount,
				CommentAmount: formCommentAmount,
				Comments:      nil,
				CreationTime:  time.Now(),
			}

			// Append created post into slice
			data = append(data, post)
		}

	case "GET":
		// Set correct headers so client can request data
		// Without correct headers there can be CORS errors etc.
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var json []byte
		var err error

		// If there is id then return specific post
		// Else return all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("GET: post with id %s\n", id)

			id, err := strconv.Atoi(id)
			if err != nil {
				return
			}

			var postExists bool

			// Find correct post by looping over all posts
			for _, v := range data {
				if v.ID == id {
					postExists = true

					json, err = json2.Marshal(v)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			if !postExists {
				logger.ErrorLogger.Printf("Post with id %d does not exist", id)
			}
		} else {
			logger.InfoLogger.Println("GET: all posts")

			if len(data) == 0 {
				logger.WarningLogger.Println("There are 0 posts")
			}

			json, err = json2.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Write json to Response
		_, err = w.Write(json)
		if err != nil {
			return
		}

	case "DELETE":
		// If there is id then delete specific post
		// Else delete all posts
		if len(id) != 0 {
			logger.InfoLogger.Printf("DELETE: post with id %s\n", id)

			id, err := strconv.Atoi(id)
			if err != nil {
				return
			}

			var sliceItemIndex int
			var indexSet bool

			for i, v := range data {
				if v.ID == id {
					sliceItemIndex = i
					indexSet = true
					break
				}
			}

			if indexSet {
				data = append(data[:sliceItemIndex], data[sliceItemIndex+1:]...)
				logger.ErrorLogger.Printf("Post with id %d deleted\n", id)
			} else {
				logger.ErrorLogger.Printf("Cannot find post with id %d\n", id)
			}

		} else {
			logger.InfoLogger.Println("DELETE: all posts")

			data = nil

			logger.InfoLogger.Println("All posts deleted")
		}
	}
}
