package app

import (
	"github.com/gentildpinto/mangope-api/config"
	"github.com/gentildpinto/mangope-api/config/database"
	"github.com/gentildpinto/mangope-api/config/logger"
	"github.com/gentildpinto/mangope-api/config/server"
)

func Start(config *config.Configuration) (err error) {
	if err = logger.Log(database.Initialize(config.Database.Db)); err != nil {
		return
	}

	e := server.New()

	initRoutes(e)

	server.Start(e, &server.Config{
		Port:         config.Server.Port,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		Debug:        config.Server.Debug,
	})

	return
}
