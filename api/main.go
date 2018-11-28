package main

import (
	"github.com/diskordanz/rest-api/api/app"
	"github.com/diskordanz/rest-api/api/config"
)

var configuration *config.Config

func main() {
	app := &app.App{}
	app.Initialize(configuration)
	app.Run(":8080")
}

func init() {
	configuration = config.GetConfig()
}
