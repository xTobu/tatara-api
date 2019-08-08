package user

import (
	"tatara-api/lib/database"
	modelUser "tatara-api/lib/database/models/user"
	"tatara-api/lib/database/repositories"
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
	result = StructDoubleUsers{
		UsersOne: users,
		UsersTwo: users,
	}
	return
}

// ReadUsers 取得所有 User
func (r *Repo) ReadUsers() (users []modelUser.User, err error) {
	err = r.DB.Find(&users).Error
	return
}

// CreateUser 新增 User
func (r *Repo) CreateUser(jsonUser *User) (err error) {
	// 使用 goroutine : 只發送，不在意儲存結果
	go func() {
	}()
	err = r.DB.Create(jsonUser).Error
	return
}

// DeleteUser 刪除 User
func (r *Repo) DeleteUser(id *string) (err error) {
	user := new(User)
	err = r.DB.Where("id = ?", &id).Delete(user).Error
	return
}

// UpdateUser 更新 User
func (r *Repo) UpdateUser(id *string, jsonUser *User) (err error) {
	user := new(User)
	err = r.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return
	}
	err = r.DB.Model(user).Updates(jsonUser).Error
	return
}
