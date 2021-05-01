package controller

import "github.com/gin-gonic/gin"

func RegisterDisplayAction(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}
