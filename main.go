package main

import (
	"tatara-api/config"
	"tatara-api/database"
	"tatara-api/server"
)

func main() {
	config.Init("development")
	database.Init()
	server.Init()
}
