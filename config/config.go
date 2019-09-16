package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "gwi-dev",
			Password: "secretpassword",
			Name:     "gwi-db",
			Charset:  "utf8",
		},
	}
}
