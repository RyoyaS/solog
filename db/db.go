package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Sex      int
	Birthday string
}

func DbInit() {
	db, err := gorm.Open("sqlite3", "solog.sqlite3")
	if err != nil {
		panic("データベースが開けません(dbInit)")
	}
	db.AutoMigrate(&User{})
	defer db.Close()
}

func DbInsert(name string, email string, password string, sex int, birthday string) {
	db, err := gorm.Open("sqlite3", "solog.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&User{Name: name, Email: email, Password: password, Sex: sex, Birthday: birthday})
	defer db.Close()
}

func DbGetOne(email string) User {
	db, err := gorm.Open("sqlite3", "solog.sqlite3")
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var user User
	db.First(&user, email)
	db.Close()
	return user
}
