package record

import (
	"1li/db"
	"1li/ent"
	"1li/ent/record"
	"1li/ent/user"
	"context"
	"fmt"
)

// Add adds a record to the database.
func Add(ctx context.Context, code string, target string, userid int) (*ent.Record, error) {
	user, err := db.Client.User.Query().
		Where(user.Tgid(userid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	rec, err := db.Client.Record.Create().
		SetCode(code).
		SetTarget(target).
		SetUser(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create record: %w", err)
	}

	return rec, nil
}

// Exists checks if a record exists in the database.
func Exists(ctx context.Context, code string) (bool, error) {
	_, err := db.Client.Record.Query().
		Where(record.Code(code)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}

		return false, fmt.Errorf("query record: %w", err)
	}

	return true, nil
}
