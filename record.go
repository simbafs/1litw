package main

import (
	"1li/ent"
	"1li/ent/user"
	"context"
)

func getUser(client *ent.Client, ctx context.Context, id int) (*ent.User, error) {
	user, err := client.User.Query().
		Where(user.Tgid(id)).
		Only(ctx)

	return user, err
}

func addRecord(client *ent.Client, ctx context.Context, code, target string, userid int) (*ent.Record, error) {
	user, err := getUser(client, ctx, userid)
	if err != nil {
		return nil, err
	}

	rec, err := client.Record.Create().
		SetCode(code).
		SetTarget(target).
		SetUser(user).
		Save(ctx)

	return rec, nil
}
