package db

import (
	"1li/db/driver"
	"1li/ent"
	"1li/ent/user"
	"context"
	"database/sql"
	"fmt"
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

func AddUser(ctx context.Context, code string, target string, userid int) (*ent.Record, error) {
	user, err := Client.User.Query().
		Where(user.Tgid(userid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	rec, err := Client.Record.Create().
		SetCode(code).
		SetTarget(target).
		SetUser(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create record: %w", err)
	}

	return rec, nil
}
