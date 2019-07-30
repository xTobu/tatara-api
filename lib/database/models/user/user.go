package user

import (
	"tatara-api/lib/database"
	"tatara-api/lib/log"
)

// User ...
type User struct {
	ID   uint64 `gorm:"type:bigserial; PRIMARY KEY; AUTO_INCREMENT:number" json:"pkid"`
	Name string `gorm:"type:varchar(20)" json:"name"`
}

// Migration ...
func Migration() {
	db := database.GetDB()
	has := db.HasTable(&User{})
	if !has {
		db.AutoMigrate(&User{})
		log.Info("AutoMigrate：user")
	}
	db.Create(&User{Name: "Junxiang"})

}

// FindUsers 取得所有 User
func FindUsers() (users []User, err error) {
	db := database.GetDB()
	err = db.Find(&users).Error
	if err != nil {
		log.Error("FindUsers", err)
		return
	}
	return
}
