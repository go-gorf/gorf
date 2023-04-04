package gorf

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type DatabaseBackend interface {
	Connect() error
	Close() error
}

type SqliteBackend struct {
	name string
}

func (backend *SqliteBackend) Connect() error {
	DB, err = gorm.Open(sqlite.Open(backend.name), &gorm.Config{})
	return err
}

func (backend *SqliteBackend) Close() error {
	return nil
}

type PostgrSQLBackend struct {
	dsn string
}

func (backend *PostgrSQLBackend) Connect() error {
	DB, err = gorm.Open(postgres.Open(backend.dsn), &gorm.Config{})
	return err
}

func (backend *PostgrSQLBackend) Close() error {
	return nil
}

func InitializeDatabase() {
	err := Settings.DbConf.Connect()
	if err != nil {
		panic("Unable to initialise the database")
	}
}
