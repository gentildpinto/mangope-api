package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gentildpinto/mangope-api/app/helpers"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	Debug        bool
}

func New() *echo.Echo {
	e := echo.New()

	requestID := middleware.RequestIDConfig{Generator: requestIDGenerator}

	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}\n"}),
		middleware.Recover(), middleware.RequestIDWithConfig(requestID),
	)

	return e
}

func Start(e *echo.Echo, config *Config) {
	srvr := &http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}

	e.Debug = config.Debug

	go func() {
		if err := e.StartServer(srvr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("server shutdown", err)
	}
}

func requestIDGenerator() string {
	return helpers.GenerateUUID()
}
