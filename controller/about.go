package controller

import "github.com/gin-gonic/gin"

func AboutDisplayAction(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{})
}
