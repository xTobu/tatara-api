package controllers

import (
	"tatara-api/server/handlers/stock"

	"github.com/gin-gonic/gin"
)

// StockInit : for individual use
func StockInit(router *gin.RouterGroup) {
	group := router.Group("stock")
	{
		group.GET("", stock.GETPrice)
	}
}
