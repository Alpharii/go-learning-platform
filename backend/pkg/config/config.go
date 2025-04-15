package config

import (
	"fmt"
	"os"
)

type Config struct {
    GoogleClientID     string
    GoogleClientSecret string
    RedirectURL        string
    DBDsn              string
}

func LoadConfig() *Config {
	RedirectURL := os.Getenv("REDIRECT_URL")
    fmt.Println("Loaded REDIRECT_URL:", RedirectURL) // Debug line
    return &Config{
        GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL:        os.Getenv("REDIRECT_URL"),
        DBDsn:              os.Getenv("DB_DSN"),
    }
}