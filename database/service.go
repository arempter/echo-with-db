package database

import (
	"context"
	"echo-with-db/errors"
	"echo-with-db/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var addUserQuery = fmt.Sprintf(`INSERT INTO %s (user_name, description) VALUES (:user_name, :description)`, usersTable)

type Database interface {
	CreateSchema() *errors.Error
	AddUser(ctx context.Context, user User) bool
	Logout()
}

type service struct {
	db *sqlx.DB
}

func New(createSchema bool) (*service, *errors.Error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	s := &service{db: db}

	if createSchema {
		err = s.CreateSchema()
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s service) CreateSchema() *errors.Error {
	const op errors.Op = "schema.create"
	_, err := s.db.Exec(schema)
	if err != nil {
		return errors.E(op, errors.Msg("failed to create database schema"), err, logrus.ErrorLevel)
	}
	return nil
}

func (s service) AddUser(ctx context.Context, user User) bool {
	const op errors.Op = "user.add"
	_, err := s.db.NamedExecContext(ctx, addUserQuery, user)
	if err != nil {
		utils.Log(errors.E(op, errors.Msg("failed to add user to db"), err, logrus.WarnLevel))
		return false
	}
	return true
}

func (s service) Logout() {
	panic("implement me")
}

func Connect() (*sqlx.DB, *errors.Error) {
	const op errors.Op = "database.connect"

	db, err := sqlx.Connect("mysql", "echo:echo@/echodb")
	if err != nil {
		return nil, errors.E(op, errors.Msg("failed to connect to database"), err, logrus.ErrorLevel)
	}

	return db, nil
}
