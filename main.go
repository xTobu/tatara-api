package main

import (
	"tatara-api/config"
	"tatara-api/lib/log"
	"tatara-api/server"
)

func main() {
	config.Init("development")
	log.Init()
	server.Init()
}
