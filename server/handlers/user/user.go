package user

import (
	"tatara-api/lib/database/repositories/user"
	"tatara-api/server/handlers"

	"github.com/gin-gonic/gin"
)

// GETUsers 取得所有使用者
func GETUsers(c *gin.Context) {
	res, err := user.AllUsers()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errors": "user.FindUsers",
		})
		return
	}
	c.JSON(200, handlers.OK(res))

}
