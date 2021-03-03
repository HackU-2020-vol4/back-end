package main

import (
	"github.com/HackU-2020-val4/back-end/db"
	"github.com/HackU-2020-val4/back-end/server"
)

func main() {
	db.Init()
	server.Init()
}
