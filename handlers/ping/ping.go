package ping

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
