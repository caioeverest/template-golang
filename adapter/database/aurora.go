package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import driver
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/config"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/logger"
)

type Aurora struct {
	client          *sql.DB
	coldStorageAddr string
}

const (
	driverAurora         = "mysql"
	auroraImplementation = "aurora"
)

func init() {
	_strategyMap[auroraImplementation] = new(Aurora)
}

func (db *Aurora) Connect(user, pass, host, dbname, coldStorageAddr string, port uint, connMaxLifetime, maxOpenConn, maxIdleConns int) error {
	var (
		sqlDB            *sql.DB
		err              error
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	)

	if sqlDB, err = sql.Open(driverAurora, connectionString); err != nil {
		return err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifetime))
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	db.client = sqlDB
	db.coldStorageAddr = coldStorageAddr
	return nil
}

func (db *Aurora) Query(template string, args ...interface{}) (result *sql.Rows, err error) {
	return db.client.Query(template, args...)
}

func (db *Aurora) Exec(template string, args ...interface{}) (result sql.Result, err error) {
	return db.client.Exec(template, args...)
}

func (db *Aurora) Shutdown(context.Context) error {
	return db.client.Close()
}

func (db *Aurora) Health() bool {
	return db.client.Ping() == nil
}
