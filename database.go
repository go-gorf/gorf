package gorf

import "errors"

type DbOperationType int

const (
	CREATE DbOperationType = iota
	GET
	UPDATE
	DELETE
)

type Query interface {
	Describe() (string, error)
	Operation() DbOperationType
}

type QuerySet interface {
	Count() (int, error)
	First() (any, error)
	Last() (any, error)
}

type Db interface {
	Get(dest interface{}, key string) error
	Filter(dest interface{}, conds ...interface{}) error
	AutoMigrate(dst ...interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Create(value interface{}) error
	GetUser(dest interface{}, id string) error
}

type DbBackend interface {
	Connect() (Db, error)
	Disconnect() error
}

var DB Db

func InitializeDatabase() error {
	var err error
	if Settings.DbBackends == nil {
		return errors.New("no database backend present")
	}

	DB, err = Settings.DbBackends.Connect()
	if err != nil {
		return errors.New("Unable to initialise the database")
	}
	return nil
}
