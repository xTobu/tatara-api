package ping

import (
	"tatara-api/lib/log"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context) {
	log.Info("Ping pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
