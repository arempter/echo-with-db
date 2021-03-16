package mysql

import (
	"echo-with-db/errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var schema = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INT AUTO_INCREMENT,    	
		user_name VARCHAR(10) NOT NULL UNIQUE,
		description VARCHAR(20) NOT NULL,
		PRIMARY KEY (id)	
	);`, UsersTable)

func CreateSchema(db *sqlx.DB) *errors.Error {
	const op errors.Op = "schema.create"
	_, err := db.Exec(schema)
	if err != nil {
		return errors.E(op, errors.Msg("failed to create database schema for mysql"), err, logrus.ErrorLevel)
	}
	return nil
}
