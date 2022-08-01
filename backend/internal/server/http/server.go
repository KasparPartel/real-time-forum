package http

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/server/http/handlers"
	config2 "real-time-forum/pkg/config"
	"real-time-forum/pkg/logger"
	"real-time-forum/pkg/websockets"
	// "real-time-forum/pkg/handlers"
)

func RunHTTPServer() {
	// Load config from config pkg
	cfg := &config2.Config

	mux := http.NewServeMux()

	// Multiplexer handlers
	mux.HandleFunc(fmt.Sprintf("/v%s/api/post/", cfg.Version), handlers.PostHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/user/", cfg.Version), handlers.UserHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/comments/", cfg.Version), handlers.CommentHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/categories/", cfg.Version), handlers.CategoryHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/like/", cfg.Version), handlers.LikeHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/dislike/", cfg.Version), handlers.DislikeHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/user/me", cfg.Version), handlers.MeHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/login/", cfg.Version), handlers.LoginHandler)
	mux.HandleFunc(fmt.Sprintf("/v%s/api/logout/", cfg.Version), handlers.LogoutHandler)
	
	// Websocket handler
	// mux.HandleFunc(fmt.Sprintf("/v%s/api/ws", cfg.Version), handlers.Upgrade)
	// mux.HandleFunc(fmt.Sprintf("/v%s/api/ws", cfg.Version), handlers.WsEndpoint)
	
    // http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	pool := websockets.NewPool()
	go pool.Start()
		
	mux.HandleFunc(fmt.Sprintf("/v%s/api/ws", cfg.Version), func(w http.ResponseWriter, r *http.Request) {
		handlers.WsEndpoint(pool, w, r)
	})
	
	// http.HandleFunc("/v%s/api/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.WsEndpoint(pool, w, r)
	// })
		
	logger.InfoLogger.Printf("Server started at http://localhost:%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux))
}
