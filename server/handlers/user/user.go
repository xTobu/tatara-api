package user

import (
	"tatara-api/lib/database/repositories/user"
	"tatara-api/server/handlers"

	"github.com/gin-gonic/gin"
)

// GETUsers 取得所有使用者
func GETUsers(c *gin.Context) {
	repo := user.NewRepo()
	res, err := repo.FindUsers()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errors": "repo.FindUsers",
		})
		return
	}
	c.JSON(200, handlers.OK(res))

}

// GETDoubleUsers 取得兩份的所有使用者
func GETDoubleUsers(c *gin.Context) {
	repo := user.NewRepo()
	res, err := repo.DoubleUsers()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errors": "repo.DoubleUsers",
		})
		return
	}
	c.JSON(200, handlers.OK(res))
}
