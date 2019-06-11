package controllers

import (
	"tatara-api/handlers/ping"

	"github.com/gin-gonic/gin"
)

func pingInit(router *gin.RouterGroup) {
	group := router.Group("ping")
	{
		group.GET("", ping.Response)
	}
}

// PingInit : for individual use
func PingInit(router *gin.RouterGroup) {
	group := router.Group("ping")
	{
		group.GET("", ping.Response)
	}
}
