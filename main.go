package main

import (
	"tatara-api/lib/config"
	"tatara-api/lib/database"
	"tatara-api/server"
)

func main() {
	config.Init("development")
	database.Init()
	server.Init()
}
