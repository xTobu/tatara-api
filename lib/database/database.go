package database

import (
	"tatara-api/lib/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init : Initial Db connection
func Init() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=tatara dbname=tatara sslmode=disable password=0000")
	if err != nil {
		log.Error("database.Init", err)
	}
	defer db.Close()
}
