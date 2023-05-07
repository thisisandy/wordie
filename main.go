package main

import (
	"wordie/core/db"
	"wordie/core/server"
)

func main() {
	db.Instance()
	server.StartServer()
}
