package core

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConf struct {
	Name string
}

func ConnectDB(db *DbConf) {
	var err error
	DB, err = gorm.Open(sqlite.Open(db.Name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func InitializeDatabase(db *DbConf) {
	ConnectDB(db)
}
