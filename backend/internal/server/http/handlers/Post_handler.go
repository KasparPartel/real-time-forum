package handlers

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/db"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homePage")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json, err := json2.Marshal(db.DummyPosts)
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(json)
	if err != nil {
		return
	}
}
