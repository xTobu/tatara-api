package controllers

import (
	"tatara-api/server/handlers/ping"

	"github.com/gin-gonic/gin"
)

// PingInit : for individual use
func PingInit(router *gin.RouterGroup) {
	group := router.Group("ping")
	{
		group.GET("", ping.Response)
	}
}