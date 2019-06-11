package main

import (
	"tatara-api/config"
	"tatara-api/server"
)
func main() {
	config.Init("development")
	server.Init()
}