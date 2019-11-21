package models

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password"  db:"password"`
	Article  []Article
}

func GetUserByUsername(username *string) (error, *User) {
	user := User{}
	err := db.Get(&user, db.Rebind("SELECT * FROM users WHERE username = ? "), username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return errors.WithStack(err), &user
	}
	return err, &user
}

func (User) TableName() string {
	return "users"
}

func AddUser(user *User) {
	Insert(db, user)
}
