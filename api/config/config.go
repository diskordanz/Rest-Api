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

	Host string
	Port string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "postgres",
			Password: "postgres",
			Name:     "books_db",
			Charset:  "utf8",

			Host: "db",
			Port: "5432",
		},
	}
}
