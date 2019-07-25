package database

import (
	"fmt"
	"tatara-api/lib/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init : Initial Db connection
func Init(sslmode, host, dbname, user, password string) {
	connArgs := fmt.Sprintf("sslmode=%s host=%s dbname=%s user=%s password=%s", sslmode, host, dbname, user, password)
	db, err := gorm.Open("postgres", connArgs)
	if err != nil {
		log.Error("database.Init", err)
	}
	defer db.Close()
}
