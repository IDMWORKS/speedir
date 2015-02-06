package datacontext

import (
	"github.com/nwoolls/speedir/errors"

	"database/sql"

	//_ to prevent golint from removing the import
	_ "github.com/lib/pq"

	"gopkg.in/gorp.v1"
)

//User model in the DB
type User struct {
	Id           int64
	Created      int64
	Username     string
	PasswordHash string
}

//InitDb creates / updates the DB schema as needed
func InitDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "user=speedir dbname=speedir sslmode=disable")
	errors.CheckErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	errors.CheckErr(err, "Create tables failed")

	return dbmap
}
