package user

import (
	"1li/db"
	"1li/ent"
	"1li/ent/user"
	"context"
)

var admin = map[int64]bool{901756183: true}

// Add adds a user to the database.
func Add(ctx context.Context, username string, tgid int64) (*ent.User, error) {
	u, err := db.Client.User.Create().
		SetTgid(tgid).
		SetUsername(username).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if isAdmin, ok := admin[tgid]; ok && isAdmin {
		return Op(ctx, username)
	}

	return u, nil
}

// Get gets a user from the database by tgid.
func Get(ctx context.Context, tgid int64) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Tgid(tgid)).Only(context.Background())
}

// GetByUsername gets a user from the database by username.
func GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
}

// IsAdmin checks if a user is an admin.
func IsAdmin(ctx context.Context, tgid int64) bool {
	u, err := Get(ctx, tgid)
	if err != nil {
		return false
	}

	return u.Admin
}

// SetAdmin sets a user as an admin.
func SetAdmin(ctx context.Context, username string, admin bool) error {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	_, err = db.Client.User.UpdateOneID(u.ID).SetAdmin(admin).Save(ctx)
	return err
}

// CanCustomCode checks if a user can use custom code.
func CanCustomCode(ctx context.Context, tgid int64) bool {
	u, err := Get(ctx, tgid)
	if err != nil {
		return false
	}

	return u.CustomCode
}

// SetCustomCode sets if a user can use custom code.
func SetCustomCode(ctx context.Context, username string, customCode bool) error {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	_, err = db.Client.User.UpdateOneID(u.ID).SetCustomCode(customCode).Save(ctx)
	return err
}

// CanReadAll checks if a user can read all records.
func CanReadAll(ctx context.Context, tgid int64) bool {
	u, err := Get(ctx, tgid)
	if err != nil {
		return false
	}
	return u.ReadAll
}

// SetReadAll sets if a user can read all records.
func SetReadAll(ctx context.Context, username string, readAll bool) error {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	_, err = db.Client.User.UpdateOneID(u.ID).SetReadAll(readAll).Save(ctx)
	return err
}

// Op sets a user as an admin.
func Op(ctx context.Context, username string) (*ent.User, error) {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return u.Update().
		SetAdmin(true).
		SetCustomCode(true).
		SetReadAll(true).
		Save(ctx)
}

// Deop sets a user as a normal user.
func Deop(ctx context.Context, username string) (*ent.User, error) {
	u, err := GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return u.Update().
		SetAdmin(false).
		SetCustomCode(false).
		SetReadAll(false).
		Save(ctx)
}
