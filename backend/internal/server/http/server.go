package http

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/server/http/handlers"
	config2 "real-time-forum/pkg/config"
	"real-time-forum/pkg/logger"
)

func RunHTTPServer() {
	// Load config from config pkg
	cfg := &config2.Config

	mux := http.NewServeMux()

	// Multiplexer handlers
	mux.HandleFunc("/api/post/", handlers.PostHandler)
	mux.HandleFunc("/api/user/", handlers.UserHandler)
	mux.HandleFunc("/api/message/", handlers.MessageHandler)
	// Websocket handler
	// mux.HandleFunc(fmt.Sprintf("/v%s/api/ws", cfg.Version), handlers.WsEndpoint)
	mux.HandleFunc("/ws", handlers.WsEndpoint)

	logger.InfoLogger.Printf("Server started at http://localhost:%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux))
}
