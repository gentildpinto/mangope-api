package controllers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

var Welcome = struct {
	Index func() echo.HandlerFunc
}{
	Index: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Welcome to Mangope-API! :) <3",
			})
		}
	},
}
