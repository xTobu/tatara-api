package server

import (
	"tatara-api/lib/config"
	"tatara-api/server/route"
)

// Init : Initial Server
func Init(c config.ServerStruct) {
	//啟動 server
	router := route.Init()
	router.Run(c.Port)
}
