package driver

import (
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
