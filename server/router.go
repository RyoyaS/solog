package server

import (
	"fmt"
	"solog/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.LoadHTMLGlob("view/*.html")
	router.Static("/static", "./static")

	router.GET("/", controller.IndexDisplayAction)
	router.GET("/about", controller.AboutDisplayAction)
	router.GET("/login", controller.LoginDisplayAction)
	router.POST("/logined", func(c *gin.Context) {
		result, user := controller.LoginUserCheck(c)
		if result {
			fmt.Println(user.Email)
		} else {
			c.HTML(200, "login.html", gin.H{"errorMessage": "メールアドレスまたはパスワードが違います"})
		}
	})
	router.GET("/register", controller.RegisterDisplayAction)
	router.GET("/registerComp", controller.RegisterCompDisplayAction)
	return router
}
