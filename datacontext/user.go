package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

// DBUser provides DB-centric methods for models.User
type DBUser struct {
	*models.User
}

// DBUsers provides DB-centric methods for a slice of DBUser
type DBUsers []*DBUser

// Scan scans each row from rows appending the populated user to users
func (users *DBUsers) scan(rows *sql.Rows) {
	for rows.Next() {
		user := &DBUser{&models.User{}}
		user.scan(rows)
		*users = append(*users, user)
	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

// Scan scans the current row in rows to populate user
func (user *DBUser) scan(rows *sql.Rows) {
	err := rows.Scan(
		&user.Id,
		&user.Created,
		&user.Username,
		&user.PasswordHash,
		&user.PasswordSalt)
	errors.CheckErr(err, "rows.Scan failed")
}

// SelectUsersByUsername returns a slicse of DBUser matching username
func (dc *DataContext) SelectUsersByUsername(username string) DBUsers {
	users := make(DBUsers, 0)

	rows, err := dc.DB.Query(sqlSelectUserByUsername, username)
	errors.CheckErr(err, "Select failed")
	users.scan(rows)

	return users
}
