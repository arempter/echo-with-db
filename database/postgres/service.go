package postgres

import (
	"echo-with-db/config"
	"echo-with-db/database"
	"echo-with-db/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type postgres struct {
	db    *sqlx.DB
	users database.UserService
	cfg   config.Config
}

func (p postgres) Users() database.UserService {
	if p.users == nil {
		p.users = NewUsers(p.db)
	}
	return p.users
}

func New(db *sqlx.DB, initSchema bool, cfg config.Config) (*postgres, *errors.Error) {
	s := &postgres{db: db, cfg: cfg}
	if initSchema {
		err := createSchema(s.db)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (p postgres) Logout() {
	panic("implement me")
}
