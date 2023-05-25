package gorf

import "errors"

type GorfDbOperationType int

const (
	CREATE GorfDbOperationType = iota
	GET
	UPDATE
	DELETE
)

type GorfQuery interface {
	Describe() (string, error)
	Operation() GorfDbOperationType
}

type GorfDbResult interface {
	Data() (string, error)
}

type GorfDB interface {
	Get(query GorfQuery) (GorfDbResult, error)
	Put(query GorfQuery) (GorfDbResult, error)
	Delete(query GorfQuery) (GorfDbResult, error)
}

type GorfDbBackend interface {
	Connect() (*GorfDB, error)
	Close() error
}

var DB *GorfDB

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
