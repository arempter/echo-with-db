package database

import "context"

type UserService interface {
	AddUser(ctx context.Context, user User) bool
}

type Database interface {
	Users() UserService
	Logout()
}

type User struct {
	id          string `db:id`
	UserName    string `db:"user_name"`
	Description string `db:"description"`
}
