package stock

import (
	"github.com/gin-gonic/gin"
)

func Price(c *gin.Context) {

	c.JSON(200, gin.H{
		"price": 10000,
	})
}
