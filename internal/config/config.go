package config

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
	Addr         string `json:"addr"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}


func LoadConfig() AppConfig{
	return AppConfig{
		Server: ServerConfig{},
		DB: DBConfig{"localhost",5432,"postgres","harsha","todo","disable"},
	}
}
