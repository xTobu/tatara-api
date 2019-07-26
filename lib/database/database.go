package database

import (
	"fmt"
	"tatara-api/lib/config"
	"tatara-api/lib/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Init : Initial Db connection
func Init(c config.DatabaseStruct) {
	connArgs := fmt.Sprintf("sslmode=%s host=%s dbname=%s user=%s password=%s", c.SSLMode, c.Host, c.DBName, c.User, c.Password)
	db, err = gorm.Open("postgres", connArgs)
	if err != nil {
		log.Error("database.Init", err)
	}
}

// GetDB : 取得 db
func GetDB() *gorm.DB {
	return db
}
