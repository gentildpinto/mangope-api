package app

import (
	"github.com/gentildpinto/mangope-api/app/controllers"
	echo "github.com/labstack/echo/v4"
)

var (
	county_controller   = controllers.County
	welcome_controller  = controllers.Welcome
	province_controller = controllers.Province
)

func initRoutes(e *echo.Echo) {
	version_one := e.Group("/api/v1")

	version_one.GET("/", welcome_controller.Index())

	// provinces
	version_one.GET("/provinces/all", province_controller.Index())
	version_one.GET("/provinces/:id", province_controller.Show())

	// counties
	version_one.GET("/counties/all", county_controller.Index())
	version_one.GET("/counties/:id", county_controller.Show())
}
