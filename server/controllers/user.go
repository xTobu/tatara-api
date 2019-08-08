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
		group.GET(":id", user.GETUser)
		// group.GET("double", user.GETDoubleUsers)
		group.POST("", user.POSTUser)
		group.DELETE(":id", user.DELETEUser)
		group.PUT(":id", user.PUTUser)
	}
}
