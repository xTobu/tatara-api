package routers

import (
	"tatara-api/controllers"
	"tatara-api/middleware"

	"github.com/gin-gonic/gin"
)

// Init :
func Init() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	{
		controllers.PingInit(api)
		controllers.StockInit(api)
		// controllers.Init(api)
	}

	auth := router.Group("auth")
	{
		auth.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"auth": true,
			})
		})

		auth.GET("mw", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"auth": true,
			})
		})
	}

	return router

}
