package database

import (
	"fmt"
	"tatara-api/lib/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init : Initial Db connection
func Init(sslmode, host, dbname, user, password string) {
	connArgs := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, sslmode, dbname, user, password)
	db, err := gorm.Open("postgres", connArgs)
	if err != nil {
		log.Error("database.Init", err)
	}
	defer db.Close()
}
