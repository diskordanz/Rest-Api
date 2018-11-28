package main

import (
	"os"
	"rest-api/api/app"
	"time"
)

var config app.Config

func main() {
	time.Sleep(10 * time.Second)
	app := &app.App{}
	app.Initialize(&config)
	app.Run(":8080")
}

func init() {
	config = app.Config{
		DB: &app.DBConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
