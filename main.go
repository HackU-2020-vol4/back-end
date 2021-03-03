package main

import (
	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/server"
)

func main() {
	db.Init()
	server.Init()
}
