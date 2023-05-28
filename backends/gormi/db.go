package gormi

import (
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func (d *GormDB) Get(dest interface{}, key string) error {
	return nil
}

func (d *GormDB) Filter(dest interface{}, conds ...interface{}) error {
	return nil
}

func (d *GormDB) AutoMigrate(dst ...interface{}) error {
	return d.DB.Migrator().AutoMigrate(dst...)
}

func (d *GormDB) First(dest interface{}, conds ...interface{}) error {
	d.DB.First(dest, conds)
	return nil
}
func (d *GormDB) Create(value interface{}) error {
	d.DB.Create(value)
	return nil
}
