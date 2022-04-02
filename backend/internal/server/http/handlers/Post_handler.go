package handlers

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/db"
	"real-time-forum/pkg/model"
	"strconv"
	"strings"
	"time"
)

var data []model.Post

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homePage")

	// Extract id from URL
	id := strings.TrimPrefix(r.URL.Path, "/post/")

	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			return
		}

		formID, err := strconv.Atoi(r.FormValue("formID"))
		if err != nil {
			fmt.Println(formID)
			fmt.Println("formID is not present or wrong format")
			return
		}
		formLikeAmount, err := strconv.Atoi(r.FormValue("formLikeAmount"))
		if err != nil {
			fmt.Println("formLikeAmount is not present or wrong format")
			return
		}
		formDislikeAmount, err := strconv.Atoi(r.FormValue("formDislikeAmount"))
		if err != nil {
			fmt.Println("formDislikeAmount is not present or wrong format")
			return
		}
		formCommentAmount, err := strconv.Atoi(r.FormValue("formCommentAmount"))
		if err != nil {
			fmt.Println("formCommentAmount is not present or wrong format")
			return
		}

		post := model.Post{
			ID:            formID,
			Title:         r.FormValue("title"),
			Body:          r.FormValue("body"),
			Author:        db.DummyUsers[0],
			Filename:      r.FormValue("image.jpg"),
			LikeAmount:    formLikeAmount,
			DislikeAmount: formDislikeAmount,
			CommentAmount: formCommentAmount,
			Comments:      nil,
			CreationTime:  time.Now(),
		}

		data = append(data, post)

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var json []byte
		var err error

		if len(id) != 0 {
			fmt.Printf("Get post with id %s\n", id)

			id, err := strconv.Atoi(id)

			for _, v := range data {
				if v.ID == id {
					json, err = json2.Marshal(v)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		} else {
			fmt.Println("Get all posts")

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
		if len(id) != 0 {
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
				fmt.Printf("Post with id %d deleted\n", id)
			} else {
				fmt.Printf("Cannot find post with id %d\n", id)
			}

		} else {
			data = nil

			fmt.Println("All posts deleted")
		}
	}
}
