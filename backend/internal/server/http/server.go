package http

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/server/http/handlers"
)

func RunHTTPServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/post/", handlers.PostHandler)

	fmt.Println("Server started at http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
