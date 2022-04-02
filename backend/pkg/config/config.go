package config

type Configuration struct {
	Port       string
	ServerName string
	DBuser     string
	DBpassword string
	DB         string
}

var Config Configuration
