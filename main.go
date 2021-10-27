package main

import (
	"log"

	"github.com/gentildpinto/mangope-api/app"
	"github.com/gentildpinto/mangope-api/config"
)

const version = "1.0.0"

func main() {
	configuration, err := config.Initialize(version)

	if err != nil {
		log.Fatal(err)
	}

	app.Start(configuration)
}
