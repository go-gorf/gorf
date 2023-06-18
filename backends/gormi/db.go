package gormi

import (
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func (d *GormDB) Get(dest interface{}, key string) error {
	//TODO:implement
	return nil
}

func (d *GormDB) GetUser(dest interface{}, id string) error {
	//TODO:implement
	return nil
}

func (d *GormDB) Filter(dest interface{}, conds ...interface{}) error {
	//TODO:implement
	return nil
}

func (d *GormDB) AutoMigrate(dst ...interface{}) error {
	//TODO:implement
	return d.DB.Migrator().AutoMigrate(dst...)
}

func (d *GormDB) First(dest interface{}, conds ...interface{}) error {
	return d.DB.First(dest, conds).Error
}
func (d *GormDB) Create(value interface{}) error {
	return d.DB.Create(value).Error
}
