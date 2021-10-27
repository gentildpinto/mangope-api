package controllers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type welcomeController struct{}

var Welcome welcomeController

func (welcomeController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to Mangope-API! :) <3",
		})
	}
}
