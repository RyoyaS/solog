package controller

import (
	"log"
	"solog/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginDisplayAction(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func Login(c *gin.Context, Email string) {
	session := sessions.Default(c)
	session.Set("Email", Email)
	session.Save()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func LoginUserCheck(c *gin.Context) (bool, db.User) {
	email, password := EmptyInputCheck(c)
	if email != "" && password != "" {
		user := db.DbUserGetOne(email)
		passMachResult := db.UserPassMach(user.Password, password)
		return passMachResult, user
	} else {
		log.Println("No value Email or Password")
		user := db.User{}
		return false, user
	}
}

func EmptyInputCheck(c *gin.Context) (string, string) {
	c.Request.ParseForm()
	email := c.Request.Form["email"]
	password := c.Request.Form["password"]
	return email[0], password[0]
}
