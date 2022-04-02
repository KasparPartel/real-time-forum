package main

import (
	"real-time-forum/internal/server/http"
	config2 "real-time-forum/pkg/config"
)

func main() {
	cfg := &config2.Config
	*cfg = config2.Configuration{
		Port:       "4000",
		ServerName: "localhost",
		DBuser:     "",
		DBpassword: "",
		DB:         "forum.db",
	}

	http.RunHTTPServer()
}
