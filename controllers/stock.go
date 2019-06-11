package controllers

import (
	"tatara-api/handlers/stock"

	"github.com/gin-gonic/gin"
)

func stockInit(router *gin.RouterGroup) {
	group := router.Group("stock")
	{
		group.GET("", stock.Price)
	}
}

// StockInit : for individual use
func StockInit(router *gin.RouterGroup) {
	group := router.Group("stock")
	{
		group.GET("", stock.Price)
	}
}
