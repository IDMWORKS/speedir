package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

type DBUser struct {
	*models.User
}

type DBUsers []*DBUser

func (users *DBUsers) Scan(rows *sql.Rows) {
	for rows.Next() {

		user := &DBUser{&models.User{}}
		user.Scan(rows)
		*users = append(*users, user)

	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

func (user *DBUser) Scan(rows *sql.Rows) {
	err := rows.Scan(
		&user.Id,
		&user.Created,
		&user.Username,
		&user.PasswordHash,
		&user.PasswordSalt)
	errors.CheckErr(err, "rows.Scan failed")
}
