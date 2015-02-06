package models

//User model in the DB
type User struct {
	Id           int64
	Created      int64
	Username     string
	PasswordHash string
}
