package api

import (
	"ciunikine/server/brand"
	"ciunikine/server/common/maps"

	"github.com/labstack/echo/v4"
)

func Brand(c echo.Context) error {
	return Success(c, maps.Map{
		"name":      brand.Name,
		"copyright": brand.Copyright,
	})
}
