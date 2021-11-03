package controllers

import (
	"net/http"
	"strconv"

	"github.com/gentildpinto/mangope-api/app/models/province"
	echo "github.com/labstack/echo/v4"
)

var Province = struct {
	Index func() echo.HandlerFunc
	Show  func() echo.HandlerFunc
}{
	Index: func() echo.HandlerFunc {
		return func(c echo.Context) error {
			provinces, err := province.GetAll()

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, provinces)
		}
	},
	Show: func() echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				return
			}

			province, err := province.GetByID(id)

			if err != nil {
				return
			}

			return c.JSON(http.StatusOK, province)
		}
	},
}
