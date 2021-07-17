package database

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
)

type SQLite struct {
	client          *sql.DB
	coldStorageAddr string
}

func StartSQLite(dbname string) (*SQLite, error) {

	db, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", dbname)) // Open the created SQLite File

	return &SQLite{db, ""}, nil
}

func (db *SQLite) Connect(string, string, string, string, string, uint, int, int, int) error {
	return nil
}

func (db *SQLite) Query(template string, args ...interface{}) (result *sql.Rows, err error) {
	return db.client.Query(template, args...)
}

func (db *SQLite) Exec(template string, args ...interface{}) (result sql.Result, err error) {
	return db.client.Exec(template, args...)
}

func (db *SQLite) ExportQueryResult(schema, reference, template string, args ...interface{}) (result sql.Result, err error) {

	result, err = db.client.Exec(template, args...)
	if err != nil {
		reError, _ := regexp.Compile(mysqlFileAlreadyExistsError)
		if reError.MatchString(err.Error()) {
			return nil, nil
		}
		return nil, err
	}

	return result, err
}

func (db *SQLite) Shutdown(context.Context) error {
	return db.client.Close()
}

func (db *SQLite) Health() bool {
	return db.client.Ping() != nil
}
