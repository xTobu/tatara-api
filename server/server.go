package server

import (
	"tatara-api/lib/config"
	"tatara-api/server/route"
)

// Init : Initial Server
func Init() {
	// 取得 port
	config := config.GetConfig()
	port := config.GetString("server.port")

	//啟動 server
	router := route.Init()
	router.Run(port)
}
