package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterConfDisplayAction(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	sex := c.PostForm("sex")
	birthday := c.PostForm("birthday")
	c.HTML(200, "registerConf.html", gin.H{"name": name, "email": email, "sex": sex, "birthday": birthday})
}
