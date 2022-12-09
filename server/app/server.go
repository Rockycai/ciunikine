package app

import (
	"ciunikine/api"

	"github.com/labstack/echo/v4"
)

func setupRoutes() *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.GET("/branding", api.Brand)

	return e
}
