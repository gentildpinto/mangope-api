package main

import (
	"log"

	"github.com/gentildpinto/mangope-api/app"
	"github.com/gentildpinto/mangope-api/config"
	"github.com/subosito/gotenv"
)

const version = "1.0.0"

func init() {
	gotenv.Load()
}

func main() {
	configuration, err := config.Initialize(version)

	if err != nil {
		log.Fatal(err)
	}

	app.Start(configuration)
}
