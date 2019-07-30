package repositories

import (
	modelUser "tatara-api/lib/database/models/user"
)

// Init : Initial repositories
func Init() {
	MigrateTables()
}

// MigrateTables : 檢查並創建所有的表
func MigrateTables() {
	modelUser.Migration()
}
