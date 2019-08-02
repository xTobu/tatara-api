package user

import (
	"tatara-api/lib/database"
	"tatara-api/lib/log"
)

// User ...
type User struct {
	ID   uint64 `json:"pkid" gorm:"type:bigserial; PRIMARY KEY; AUTO_INCREMENT:number;"` // Permiry Key
	Name string `json:"name" gorm:"type:varchar(20)"`                                    // 姓名
}

// Migration ...
func Migration() {
	db := database.GetDB()
	has := db.HasTable(&User{})
	if !has {
		db.AutoMigrate(&User{})
		log.Info("AutoMigrate：user")
	}
	// db.Create(&User{Name: "Junxiang"})
}

// AfterFind : GORM Hooks
// process after user find
// ID 為奇數時， Name 改為 "Huang"
func (user *User) AfterFind() (err error) {
	switch user.ID % 2 {
	case 1:
		user.Name = "Huang"
	}
	return
}
