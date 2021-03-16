package mysql

import (
	"echo-with-db/config"
	"echo-with-db/database"
	"echo-with-db/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mysql struct {
	db    *sqlx.DB
	users database.UserService
	cfg   config.Config
}

func (s mysql) Users() database.UserService {
	if s.users == nil {
		s.users = NewUsers(s.db)
	}
	return s.users
}

func New(db *sqlx.DB, initSchema bool, cfg config.Config) (*mysql, *errors.Error) {
	s := &mysql{db: db, cfg: cfg}
	if initSchema {
		err := createSchema(s.db)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s mysql) Logout() {
	panic("implement me")
}
