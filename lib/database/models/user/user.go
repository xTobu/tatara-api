package user

import (
	"tatara-api/lib/database"
	"tatara-api/lib/log"
	"time"
)

// User ...
type User struct {
	ID        uint64    `json:"-" gorm:"type:bigserial; PRIMARY KEY; AUTO_INCREMENT:number;"` // Permiry Key
	Name      string    `json:"name,omitempty" gorm:"column:name; type:varchar(20)"`          // 姓名
	Sex       string    `json:"sex,omitempty" gorm:"column:sex; type:varchar(20)"`            // 性別
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp(6) without time zone; not null;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp(6) without time zone; not null;"`
}

// Migration ...
func Migration() {
	db := database.GetDB()
	has := db.HasTable(&User{})
	if !has {
		db.AutoMigrate(&User{})
		log.Info("AutoMigrate：user")
		// db.Create(&User{Name: "Junxiang", Sex: "Man"})
	}
}

// AfterFind : GORM Hooks
// process after user find
// ID 為奇數時， Name 改為 "Huang"
func (user *User) AfterFind() (err error) {
	// switch user.ID % 2 {
	// case 1:
	// 	user.Name = "Huang"
	// }
	return
}
