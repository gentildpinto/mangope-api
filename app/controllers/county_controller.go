package controllers

import (
	"net/http"
	"strconv"

	"github.com/gentildpinto/mangope-api/app/models/county"
	"github.com/labstack/echo/v4"
)

var County = struct {
	Index func() echo.HandlerFunc
	Show  func() echo.HandlerFunc
}{
	Index: func() echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			counties, err := county.GetAll()

			if err != nil {
				return
			}

			return c.JSON(http.StatusOK, counties)
		}
	},
	Show: func() echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				return
			}

			county, err := county.GetByID(id)

			if err != nil {
				return
			}

			return c.JSON(http.StatusOK, county)
		}
	},
}
