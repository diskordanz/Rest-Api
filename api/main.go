package main

import (
	"os"
	"time"

	app "github.com/diskordanz/rest-api/api/app"

	_ "github.com/lib/pq"
)

var appl *app.App

func main() {
	appl.Run(":8080")
}

func init() {
	time.Sleep(10 * time.Second)

	appl.Initialize(
		&app.Config{
			DB: &app.DBConfig{
				Dialect:  os.Getenv("DB_DIALECT"),
				Username: os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				Name:     os.Getenv("DB_NAME"),
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
			},
		})
}
