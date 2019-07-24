package main

import (
	"tatara-api/lib/config"
	"tatara-api/lib/database"
	"tatara-api/server"
)

func main() {
	config.Init("development")

	// 取得 port
	config := config.GetConfig()
	sslmode := config.GetString("server.sslmode")
	host := config.GetString("server.host")
	dbname := config.GetString("server.dbname")
	user := config.GetString("server.user")
	password := config.GetString("server.password")
	database.Init(sslmode, host, dbname, user, password)

	server.Init()
}
