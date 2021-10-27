package config

import (
	"fmt"

	"github.com/gentildpinto/mangope-api/app/helpers"
	"github.com/gentildpinto/mangope-api/config/logger"
	"github.com/gentildpinto/mangope-api/config/orm"
	"gorm.io/gorm"
)

var (
	version string
	appPort = "80"
	envVars = []string{
		"DATABASE_USERNAME", "DATABASE_PASSWORD", "ENVIRONMENT", "DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME",
	}
)

type (
	Configuration struct {
		Server   *Server
		Database *Database
	}

	Server struct {
		Port              string
		Debug             bool
		ReadTimeout       int
		WriteTimeout      int
		RequestsPerSecond int
	}

	Database struct {
		Host         string
		Port         string
		User         string
		Password     string
		DatabaseName string
		LogQueries   bool
		Db           *gorm.DB
	}
)

func Initialize(appVersion string) (config *Configuration, err error) {
	version = appVersion

	err = logger.Initialize(version)

	if err != nil {
		return
	}

	if err = validateEnvironment(); err != nil {
		return
	}

	if helpers.ViperEnvVariable("APP_PORT") != "" {
		appPort = helpers.ViperEnvVariable("APP_PORT")
	}

	config = &Configuration{
		Server: &Server{
			Port:         appPort,
			Debug:        helpers.ViperEnvVariable("ENVIRONMENT") != "production" || helpers.ViperEnvVariable("DEBUG") == "true",
			ReadTimeout:  60,
			WriteTimeout: 60,
		},
		Database: &Database{
			Host:         helpers.ViperEnvVariable("DATABASE_HOST"),
			Port:         helpers.ViperEnvVariable("DATABASE_PORT"),
			User:         helpers.ViperEnvVariable("DATABASE_USERNAME"),
			Password:     helpers.ViperEnvVariable("DATABASE_PASSWORD"),
			DatabaseName: helpers.ViperEnvVariable("DATABASE_NAME"),
			LogQueries:   helpers.ViperEnvVariable("ENVIRONMENT") != "production" || helpers.ViperEnvVariable("DEBUG") == "true",
		},
	}

	config.Database.Db, err = orm.New(config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName)

	if err = logger.Log(err); err != nil {
		return
	}

	return
}

func validateEnvironment() error {
	for _, envVar := range envVars {
		if helpers.ViperEnvVariable(envVar) == "" {
			return logger.Log(fmt.Errorf("missing environment variable: %s", envVar))
		}
	}
	return nil
}
