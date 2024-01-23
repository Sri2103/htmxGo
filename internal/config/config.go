package config

import (
	"os"

	"github.com/alexedwards/scs/v2"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SslMode  string `json:"sslMode"`
}

type AppConfig struct {
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"database"`
}

type ServerConfig struct {
	Addr           string `json:"addr"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	SessionManager *scs.SessionManager
}

func LoadConfig() *AppConfig {

	// viper.AutomaticEnv()
	// viper.ReadConfig()

	return &AppConfig{
		Server: ServerConfig{
			Addr: ":3500",
		},
		DB: DBConfig{

			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     5432,
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DBName"),
			SslMode:  os.Getenv("POSTGRES_SSLMode"),
		},
	}
}
