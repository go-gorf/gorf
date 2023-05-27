package gormi

import "github.com/go-gorf/gorf"

type GorfDbResult interface {
	Data() (string, error)
}

type GormDB struct {
}

func (d *GormDB) Get(query gorf.GorfQuery) (gorf.GorfDbResult, error) {
	return nil, nil
}

func (d *GormDB) Put(query gorf.GorfQuery) (gorf.GorfDbResult, error) {
	return nil, nil
}

func (d *GormDB) Delete(query gorf.GorfQuery) (gorf.GorfDbResult, error) {
	return nil, nil
}

type GormSqliteBackend struct {
}

func (b *GormSqliteBackend) Connect() (*gorf.GorfDB, error) {
	return nil, nil
}
func (b *GormSqliteBackend) Close() error {
	return nil
}
