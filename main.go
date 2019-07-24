package main

import (
	"tatara-api/database"
	"tatara-api/lib/config"
	"tatara-api/server"
)

func main() {
	config.Init("development")
	database.Init()
	server.Init()
}
