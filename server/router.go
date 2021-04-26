package server

import (
	"solog/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.Static("/static", "./static")

	router.GET("/", controller.IndexDisplayAction)
	router.GET("/about", controller.AboutDisplayAction)

	return router
}
