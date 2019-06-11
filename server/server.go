package server

import (
	"tatara-api/config"
	"tatara-api/server/routers"
)

func Init() {
	config := config.GetConfig()
	r := routers.Init()
	r.Run(config.GetString("server.port"))
}
