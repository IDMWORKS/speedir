package datacontext

import (
	"database/sql"
	"fmt"

	"github.com/idmworks/speedir/models"
)

// DBUser provides DB-centric methods for models.User
type DBUser struct {
	*models.User
}

// DBUsers provides DB-centric methods for a slice of DBUser
type DBUsers []*DBUser

// Scan scans each row from rows appending the populated user to users
func (users *DBUsers) scan(rows *sql.Rows) error {
	for rows.Next() {
		user := &DBUser{&models.User{}}
		user.scan(rows)
		*users = append(*users, user)
	}
	return rows.Err()
}

// Scan scans the current row in rows to populate user
func (user *DBUser) scan(rows *sql.Rows) error {
	err := rows.Scan(
		&user.Id,
		&user.Created,
		&user.Username,
		&user.PasswordHash,
		&user.PasswordSalt)
	return err
}

// SelectUsersByUsername returns a slice of DBUser matching username
func (dc *DataContext) SelectUsersByUsername(username string) (result DBUsers, err error) {
	users := make(DBUsers, 0)

	rows, err := dc.DB.Query(sqlSelectUserByUsername, username)
	if err != nil {
		return nil, fmt.Errorf("SelectUsersByUsername failed: %v", err)
	}

	users.scan(rows)
	return users, nil
}
