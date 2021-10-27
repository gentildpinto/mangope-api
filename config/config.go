package config

import (
	"fmt"
	"os"

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
		SSLMode      string
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

	if os.Getenv("APP_PORT") != "" {
		appPort = os.Getenv("APP_PORT")
	}

	config = &Configuration{
		Server: &Server{
			Port:         appPort,
			Debug:        os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true",
			ReadTimeout:  60,
			WriteTimeout: 60,
		},
		Database: &Database{
			Host:         os.Getenv("DATABASE_HOST"),
			Port:         os.Getenv("DATABASE_PORT"),
			User:         os.Getenv("DATABASE_USERNAME"),
			Password:     os.Getenv("DATABASE_PASSWORD"),
			DatabaseName: os.Getenv("DATABASE_NAME"),
			LogQueries:   os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true",
			SSLMode:      os.Getenv("SSL_MODE"),
		},
	}

	config.Database.Db, err = orm.New(config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName, config.Database.SSLMode)

	if err = logger.Log(err); err != nil {
		return
	}

	return
}

func validateEnvironment() error {
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return logger.Log(fmt.Errorf("missing environment variable: %s", envVar))
		}
	}
	return nil
}
