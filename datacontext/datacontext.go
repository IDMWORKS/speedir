package datacontext

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq" //_ = imported for side effects
	"github.com/nwoolls/speedir/errors"
	"github.com/nwoolls/speedir/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v1"
)

const (
	adminUsername = "admin"
	adminPassword = "admin"
)

//InitDb creates / updates the DB schema as needed
func InitDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "user=speedir dbname=speedir sslmode=disable")
	errors.CheckErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	errors.CheckErr(err, "Create tables failed")

	return dbmap
}

//SeedDb seeds the DB with data necessary for the app to run
func SeedDb(dbmap *gorp.DbMap) {
	count, err := dbmap.SelectInt("select count(id) from users where username=$1", adminUsername)
	errors.CheckErr(err, "SelectInt failed")

	if count == 0 {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		errors.CheckErr(err, "bcrypt.GenerateFromPassword failed")

		admin := models.User{
			Created:      time.Now().UnixNano(),
			Username:     adminUsername,
			PasswordHash: string(passwordHash),
		}

		err = dbmap.Insert(&admin)
		errors.CheckErr(err, "Insert failed")
	}
}
