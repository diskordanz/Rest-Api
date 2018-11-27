package main

import (
	"Rest-Api/api/app"
	"Rest-Api/api/config"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
