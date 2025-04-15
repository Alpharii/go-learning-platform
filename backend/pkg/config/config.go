package config

import "os"

type Config struct{
	GoogleClientID		string
	GoogleClientSecret	string
	RedirectURL			string
	DBDsn				string
}

func LoadConfig() *Config {
	return &Config{
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL: os.Getenv("REDIRECT_URL"),
		DBDsn: os.Getenv("DB_DSN"),
	}
}