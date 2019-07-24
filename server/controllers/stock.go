package controllers

import (
	"tatara-api/handlers/stock"

	"github.com/gin-gonic/gin"
)

// StockInit : for individual use
func StockInit(router *gin.RouterGroup) {
	group := router.Group("stock")
	{
		group.GET("", stock.Price)
	}
}
