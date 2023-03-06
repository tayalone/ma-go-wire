package config

import "os"

// Config ENV
type Config struct {
	Port string
}

// Initialize ENV Configuration
func Initialize() Config {
	port := os.Getenv("PORT")
	if port != "" {
		port = "3000"
	}

	return Config{
		Port: port,
	}
}
