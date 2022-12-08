package server

import (
	v1 "ciunikine/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.SetTrustedProxies([]string{"localhost"})

	home := v1.NewHomeWorkers()

	router.GET("/index", home.IndexHome)

	return router
}
