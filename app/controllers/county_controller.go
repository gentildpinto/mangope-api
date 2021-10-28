package controllers

import (
	"net/http"

	"github.com/gentildpinto/mangope-api/app/models/county"
	echo "github.com/labstack/echo/v4"
)

type countyController struct{}

var County countyController

func (countyController) Index() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		counties, err := county.GetAll()

		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"counties": counties,
		})
	}
}
