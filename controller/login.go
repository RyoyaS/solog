package controller

import "github.com/gin-gonic/gin"

func LoginDisplayAction(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}
