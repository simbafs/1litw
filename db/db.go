package db

import (
	"1li/db/driver"
	"1li/ent"
	"context"
	"database/sql"
)

var Client *ent.Client

func InitDB() error {
	// register modernc.org/sqlite as sqlite3 in database/sql
	sql.Register("sqlite3", driver.NewSqlite3Driver())

	client, err := ent.Open("sqlite3", "./1litw.sqlite?_pragma=foreign_keys(1)")
	if err != nil {
		return err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}

	Client = client

	return nil
}
