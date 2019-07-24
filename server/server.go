package server

import (
	"tatara-api/lib/config"
	"tatara-api/server/routers"
)

// Init : Initial Server
func Init() {
	config := config.GetConfig()
	r := routers.Init()
	r.Run(config.GetString("server.port"))
}
