package mysql

import (
	"context"
	"echo-with-db/database"
	"echo-with-db/errors"
	"echo-with-db/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) *users {
	return &users{db: db}
}

func (s users) AddUser(ctx context.Context, user database.User) bool {
	var addUserQuery = fmt.Sprintf(`INSERT INTO %s (user_name, description) VALUES (:user_name, :description)`, UsersTable)
	const op errors.Op = "user.add"
	_, err := s.db.NamedExecContext(ctx, addUserQuery, user)
	if err != nil {
		utils.Log(errors.E(op, errors.Msg("failed to add user to db"), err, logrus.WarnLevel))
		return false
	}
	return true
}
