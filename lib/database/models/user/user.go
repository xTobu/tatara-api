package user

import (
	"tatara-api/lib/database"
	"tatara-api/lib/log"
)

// User ...
type User struct {
	ID   string `gorm:"type:int" json:"pkid"`
	Name string `gorm:"type:varchar(20)" json:"name"`
}

// FindUsers 取得所有 User
func FindUsers() (users []*User, err error) {
	db := database.GetDB()
	err = db.Find(&users).Error
	if err != nil {
		log.Error("FindUsers", err)
		return
	}
	return
}
