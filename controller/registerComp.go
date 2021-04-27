package controller

import "github.com/gin-gonic/gin"

func RegisterCompDisplayAction(c *gin.Context) {
	c.HTML(200, "registerComp.html", gin.H{})
}
