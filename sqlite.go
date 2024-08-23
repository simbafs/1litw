package main

// https://github.com/ent/ent/issues/2460#issuecomment-1095210972

import (
	"database/sql"
	"database/sql/driver"

	"modernc.org/sqlite"
)

type Sqlite3Driver struct {
	*sqlite.Driver
}

type Sqlite3DriverConn interface {
	Exec(string, []driver.Value) (driver.Result, error)
}

func (d *Sqlite3Driver) Open(name string) (driver.Conn, error) {
	return d.Driver.Open(name)
}

func NewSqlite3Driver() *Sqlite3Driver {
	return &Sqlite3Driver{Driver: &sqlite.Driver{}}
}

// RegisterSqlite3Driver register modernc.org/sqlite as sqlite3 in database/sql
func RegisterSqlite3Driver() {
	sql.Register("sqlite3", NewSqlite3Driver())
}
