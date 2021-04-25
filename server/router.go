package server

import (
	"solog/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/", controller.IndexDisplayAction)

	return router
}
