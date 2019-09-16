package main

import (
	"log"

	"./app"
	"./config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)

	log.Fatal(app.Run())

}
