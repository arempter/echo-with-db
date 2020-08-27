package database

import "fmt"

const usersTable = "users"

var schema = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INT AUTO_INCREMENT,    	
		user_name VARCHAR(10) NOT NULL UNIQUE,
		description VARCHAR(20) NOT NULL,
		PRIMARY KEY (id)	
	);`, usersTable)

type User struct {
	id          string `db:id`
	UserName    string `db:"user_name"`
	Description string `db:"description"`
}
