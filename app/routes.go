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
	e.GET("/", welcome_controller.Index())

	// provinces
	e.GET("/provinces/all", province_controller.Index())
	e.GET("/provinces/:id", province_controller.Show())

	// counties
	e.GET("/counties/all", county_controller.Index())
	e.GET("/counties/:id", county_controller.Show())
}
