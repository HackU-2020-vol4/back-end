package main

import (
	"go-template/db"
	"go-template/server"
)

func main() {
	db.Init()
	server.Init()
}
