package user

import (
	repoUser "tatara-api/lib/database/repositories/user"
	"tatara-api/lib/log"
	"tatara-api/server/handlers"

	"github.com/gin-gonic/gin"
)

// POSTUser 新增使用者
func POSTUser(c *gin.Context) {
	req := new(repoUser.User)
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	repo := repoUser.NewRepo()
	err = repo.CreateUser(req)
	if err != nil {
		log.Error("repo.POSTUser", err)
		c.AbortWithStatusJSON(400, gin.H{
			"errors": "repo.POSTUser",
		})
		return
	}
	c.JSON(200, handlers.OK(""))

}

// GETUsers 取得所有使用者
func GETUsers(c *gin.Context) {
	repo := repoUser.NewRepo()
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
	repo := repoUser.NewRepo()
	res, err := repo.DoubleUsers()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errors": "repo.DoubleUsers",
		})
		return
	}
	c.JSON(200, handlers.OK(res))
}
