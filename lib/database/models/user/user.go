package user

import (
	"tatara-api/lib/database"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	// ID        uint64    `json:"-" gorm:"type:bigserial; PRIMARY KEY; AUTO_INCREMENT:number;"` // Permiry Key
	// CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp(6)"`
	// UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp(6)"`
	// DeletedAt time.Time `json:"deletedAt" gorm:"type:timestamp(6)"`

	Account  string `json:"account" gorm:"type:varchar(20)" binding:"required"`  // 帳號
	Password string `json:"password" gorm:"type:varchar(20)" binding:"required"` // 密碼

}

// Migration ...
func Migration() {
	db := database.GetDB()
	// HasTable 無法判斷有 schema 的 table ，
	// 故都會回傳 false
	// ref. https://github.com/jinzhu/gorm/issues/1197
	has := db.HasTable(&User{})
	if !has {
		// // Set DefaultTableName
		// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// 	return "ro." + defaultTableName
		// }
		db.AutoMigrate(&User{})
		//  log.Info("AutoMigrate：ro.user")
	}
}

// TableName ...
func (user *User) TableName() string {
	return "ro.user"
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
