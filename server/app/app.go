package app

import (
	"ciunikine/server/brand"
	"fmt"

	"github.com/labstack/echo/v4"
)

var app *App

type App struct {
	Server *echo.Echo
}

func newApp() *App {
	return &App{}
}

func init() {
	app = newApp()
}

func Run() error {
	fmt.Println(brand.Hi)

	app.Server = setupRoutes()

	app.Server.Start(":8080")

	return nil
}
