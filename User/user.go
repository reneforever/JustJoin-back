package User

import (
	"database/sql"
)

type User struct {
	ID       int8   `json: id`
	Username string `json: username`
	Password string `json: password`
}

func QueryUserInfo(db *sql.DB, username string) (userInfo *User, err error) {
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	var u1 User
	sqlStr := "select id, username, password from user where username=?"
	queryErr := db.QueryRow(sqlStr, username).Scan(&u1.ID, &u1.Username, &u1.Password)
	if queryErr != nil {
		return nil, queryErr
	}
	return &u1, nil
}
