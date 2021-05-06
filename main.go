package main

import (
	"solog/db"
	"solog/server"
)

func main() {
	db.DbInit()
	r := server.GetRouter()
	r.Run(":8080")
}
