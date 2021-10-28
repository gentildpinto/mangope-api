package app

import (
	"github.com/gentildpinto/mangope-api/app/controllers"
	echo "github.com/labstack/echo/v4"
)

var (
	welcome_controller  = controllers.Welcome
	province_controller = controllers.Province
)

func initRoutes(e *echo.Echo) {
	e.GET("/", welcome_controller.Index())
	e.GET("/provinces/all", province_controller.Index())
}
