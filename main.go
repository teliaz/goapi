package main

import (
	"log"

	"github.com/teliaz/goapi/app"
	"github.com/teliaz/goapi/config"
)

func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	log.Fatal(app.Run(":8080")) // In case this port is used elsewhere

}
