package database

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import driver
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/config"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/logger"
)

type MySQL struct {
	client          *sql.DB
	coldStorageAddr string
}

const (
	driverMySQL                 = "mysql"
	mysqlImplementation         = "mysql"
	mysqlFileAlreadyExistsError = `Error 1086: File '.*?' already exists`
)

func init() {
	_strategyMap[mysqlImplementation] = new(MySQL)
}

func (db *MySQL) Connect(user, pass, host, dbname, coldStorageAddr string, port uint, connMaxLifetime, maxOpenConn, maxIdleConns int) error {
	var (
		sqlDB            *sql.DB
		err              error
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	)

	if sqlDB, err = sql.Open(driverMySQL, connectionString); err != nil {
		return err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifetime))
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	db.client = sqlDB
	db.coldStorageAddr = coldStorageAddr
	return nil
}

func (db *MySQL) Query(template string, args ...interface{}) (result *sql.Rows, err error) {
	return db.client.Query(template, args...)
}

func (db *MySQL) Exec(template string, args ...interface{}) (result sql.Result, err error) {
	return db.client.Exec(template, args...)
}

func (db *MySQL) ExportQueryResult(schema, reference, template string, args ...interface{}) (result sql.Result, err error) {
	moveTemplate := fmt.Sprintf("%s INTO OUTFILE '%s/%s-%s.csv' FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '\"' LINES TERMINATED BY '\\n'", template, db.coldStorageAddr, reference, schema)
	logger.Get(config.Get()).Debugf("Query: %s", moveTemplate)

	result, err = db.client.Exec(moveTemplate, args...)
	if err != nil {
		reError, _ := regexp.Compile(mysqlFileAlreadyExistsError)
		if reError.MatchString(err.Error()) {
			return nil, nil
		}
		return nil, err
	}

	return result, err
}

func (db *MySQL) Shutdown(context.Context) error {
	return db.client.Close()
}

func (db *MySQL) Health() bool {
	return db.client.Ping() == nil
}
