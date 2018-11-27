package main

import (
	"rest-api/api/app"
	"rest-api/api/config"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
