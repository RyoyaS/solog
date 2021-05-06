package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoginDisplayAction(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func LoginedAction(c *gin.Context) {
	email, password := EmptyInputCheck(c)
	if email != "" && password != "" {
		//Login(email, password)
	} else {
		log.Fatal("No value Email or Password")
	}
}

func EmptyInputCheck(c *gin.Context) (string, string) {
	c.Request.ParseForm()
	email := c.Request.Form["email"]
	password := c.Request.Form["password"]

	if email[0] == "" || password[0] == "" {
		c.HTML(200, "login.html", gin.H{"aboutEmailMessage": "値が入力されていません"})
	} else {
		return email[0], password[0]
	}
	return "", ""
}

/* func Login(email, password string) {
	var hashStr = ""
	start := time.Now()
	for _, user := range users {
		if email == user.LoginId {
			hashStr = user.Password
			break
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(password))
	end := time.Now()
	fmt.Printf("%fs\t", (end.Sub(start)).Seconds())
	if err == nil {
		// 成功
		fmt.Print("Success")
	} else {
		// 失敗
		fmt.Print("Failure")
	}
	fmt.Printf("\t%s/%s\n", email, password)
} */
