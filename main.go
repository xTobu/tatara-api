package main

import (
	"tatara-api/lib/config"
	"tatara-api/lib/database"
	"tatara-api/lib/database/repositories"
	"tatara-api/server"
)

func main() {
	// 初始化 config
	config.Init("development")
	conf := config.GetConfig()

	// 初始化 database
	database.Init(conf.Database)
	repositories.Init()

	// 初始化 server
	server.Init(conf.Server)
}
