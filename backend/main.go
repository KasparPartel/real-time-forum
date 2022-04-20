package main

import (
	"real-time-forum/internal/server/http"
	config2 "real-time-forum/pkg/config"
)

func init() {
	// Populate pointer to config so it is project-wide accessible
	cfg := &config2.Config
	*cfg = config2.Configuration{
		Port:       "4000",
		ServerName: "localhost",
		DBuser:     "",
		DBpassword: "",
		DBfilename: "db/forum.db",
	}
}

func main() {
	http.RunHTTPServer()
}
