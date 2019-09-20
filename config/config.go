package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			Username: "gwiuser",
			Password: "secretpassword",
			Name:     "gwi",
			Charset:  "utf8",
		},
	}
}
