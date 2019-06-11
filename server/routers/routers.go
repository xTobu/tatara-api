package routers

import (
	"tatara-api/controllers"
	"tatara-api/controllers/stock"

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

	auth := router.Group("auth", middleware.AuthMiddleware())
	{
		auth.GET("", stock.Status)
	}

	return router

}
