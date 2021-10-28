package controllers

import (
	"net/http"

	"github.com/gentildpinto/mangope-api/app/models/province"
	echo "github.com/labstack/echo/v4"
)

type provinceController struct{}

var Province provinceController

func (provinceController) Index() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		provinces, err := province.GetAll()

		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"provinces": provinces,
		})
	}
}
