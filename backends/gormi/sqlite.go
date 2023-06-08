package gormi

import (
	"github.com/go-gorf/gorf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormSqliteBackend struct {
	Name string
}

func (b *GormSqliteBackend) Connect() (gorf.Db, error) {
	db, err := gorm.Open(sqlite.Open(b.Name), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	grfDb := &GormDB{
		DB: db,
	}
	return grfDb, nil
}

func (b *GormSqliteBackend) Disconnect() error {
	return nil
}
