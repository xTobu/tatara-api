package user

import (
	repoUser "tatara-api/lib/database/repositories/user"
	"tatara-api/lib/log"
	"tatara-api/server/handlers"

	"github.com/gin-gonic/gin"
)

// POSTUser 新增使用者
func POSTUser(c *gin.Context) {
	jsonUser := new(repoUser.User)
	err := c.BindJSON(jsonUser)
	if err != nil {
		log.Error("handler.POSTUser.BindJSON", err)
		c.AbortWithStatus(400)
		return
	}

	repo := repoUser.NewRepo()
	err = repo.CreateUser(jsonUser)
	if err != nil {
		log.Error("handler.POSTUser.CreateUser", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(true))

}

// GETUsers 取得所有使用者
func GETUsers(c *gin.Context) {
	repo := repoUser.NewRepo()
	res, err := repo.ReadUsers()
	if err != nil {
		log.Error("handler.GETUsers", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(res))

}

// GETDoubleUsers 取得兩份的所有使用者
func GETDoubleUsers(c *gin.Context) {
	repo := repoUser.NewRepo()
	res, err := repo.DoubleReadUsers()
	if err != nil {
		log.Error("handler.GETDoubleUsers", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(res))
}
