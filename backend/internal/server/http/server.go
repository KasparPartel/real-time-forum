package http

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/server/http/handlers"
	config2 "real-time-forum/pkg/config"
)

func RunHTTPServer() {
	cfg := &config2.Config

	mux := http.NewServeMux()

	mux.HandleFunc("/post/", handlers.PostHandler)

	fmt.Println("Server started at http://localhost:4000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux))
}
