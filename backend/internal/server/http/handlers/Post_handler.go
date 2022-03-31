package handlers

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/db"
	"strconv"
	"strings"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homePage")

	switch r.Method {
	case "POST":

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Extract id from URL
		id := strings.TrimPrefix(r.URL.Path, "/post/")

		var json []byte
		var err error

		if len(id) != 0 {
			fmt.Println("Get specific post")
			fmt.Println(id)

			id, err := strconv.Atoi(id)

			json, err = json2.Marshal(db.DummyPosts[id-1])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Get all posts")

			json, err = json2.Marshal(db.DummyPosts)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Write json to Response
		_, err = w.Write(json)
		if err != nil {
			return
		}
	}
}
