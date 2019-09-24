package main

import (
	"github.com/teliaz/goapi/app"
	"github.com/teliaz/goapi/config"
)

func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
