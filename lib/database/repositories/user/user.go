package user

import (
	"tatara-api/lib/database"
	modelUser "tatara-api/lib/database/models/user"
	"tatara-api/lib/database/repositories"
	"tatara-api/lib/log"
)

// Repo user repository
type Repo struct {
	repositories.Setup
}

// StructDoubleUsers ...
type StructDoubleUsers struct {
	UsersOne []modelUser.User `json:"users_one"`
	UsersTwo []modelUser.User `json:"users_two"`
}

// User ...
type User struct {
	modelUser.User
}

// NewRepo new instance of user repository
func NewRepo() *Repo {
	// 踩坑： invalid memory address
	// https://studygolang.com/articles/3174
	var r = new(Repo)
	r.DB = database.GetDB()
	return r
}

// DoubleReadUsers 結構為兩份的所有使用者
func (r *Repo) DoubleReadUsers() (result StructDoubleUsers, err error) {
	var users []modelUser.User
	err = r.DB.Find(&users).Error
	if err != nil {
		log.Error("Repo.user AllUsers", err)
		return
	}
	result = StructDoubleUsers{
		UsersOne: users,
		UsersTwo: users,
	}
	return
}

// ReadUsers 取得所有 User
func (r *Repo) ReadUsers() (users []modelUser.User, err error) {
	err = r.DB.Find(&users).Error
	if err != nil {
		log.Error("Repo.user ReadUsers", err)
		return
	}
	return
}

// CreateUser 新增 User
func (r *Repo) CreateUser(user *User) (err error) {
	// goroutine 只發送，不在意儲存結果
	go func() {
		err = r.DB.Create(user).Error
		if err != nil {
			log.Error("Repo.user CreateUser", err)
			return
		}
	}()
	return
}
