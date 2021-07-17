package database

import (
	"context"
	"database/sql"
	"errors"
)

type Database interface {
	Connect(user, pass, host, dbname, coldStorageAddr string, port uint, connMaxLifetime, maxOpenConn, maxIdleConns int) error
	Query(template string, args ...interface{}) (result *sql.Rows, err error)
	Exec(template string, args ...interface{}) (result sql.Result, err error)
	Shutdown(context context.Context) error
	Health() bool
}

var _strategyMap = make(map[string]Database)

func Connect(implementation, user, pass, host, dbname, coldStorageAddr string, port uint, connMaxLifetime, maxOpenConn, maxIdleConns int) (Database, error) {
	dbImplementation, ok := _strategyMap[implementation]
	if !ok {
		return nil, errors.New("unknown implementation")
	}

	if err := dbImplementation.Connect(user, pass, host, dbname, coldStorageAddr, port, connMaxLifetime, maxOpenConn, maxIdleConns); err != nil {
		return nil, err
	}
	return dbImplementation, nil
}
