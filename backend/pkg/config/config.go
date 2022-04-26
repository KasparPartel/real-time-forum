package config

type Configuration struct {
	Port       string
	ServerName string
	DBuser     string
	DBpassword string
	DBfilename string
	Version    string
}

// Config can be accessible from all parts of the application
var Config Configuration
