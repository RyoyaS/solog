package main

import "solog/server"

func main() {
	r := server.GetRouter()
	r.Run(":8080")
}
