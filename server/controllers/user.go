package controllers

import (
	"tatara-api/server/handlers/user"

	"github.com/gin-gonic/gin"
)

// UserInit : for individual use
func UserInit(router *gin.RouterGroup) {
	group := router.Group("user")
	{
		group.GET("", user.GETUsers)
	}
}
