package main

import (
	"back-end/db"
	"back-end/server"
)

func main() {
	db.Init()
	server.Init()
}
