package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
	_ "github.com/lib/pq" //_ = imported for side effects
	"gopkg.in/gorp.v1"
)

const (
	adminUsername = "admin"
	adminPassword = "admin"
)

// InitDb creates / updates the DB schema as needed
func InitDb() *gorp.DbMap {
	// open DB connection
	db, err := sql.Open("postgres", "user=speedir dbname=speedir sslmode=disable")
	errors.CheckErr(err, "sql.Open failed")

	// initialize gorp DB map
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// map models to tables
	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")

	// create missing tables
	err = dbmap.CreateTablesIfNotExists()
	errors.CheckErr(err, "Create tables failed")

	return dbmap
}

// SeedDb seeds the DB with data necessary for the app to run
func SeedDb(dbmap *gorp.DbMap) {
	createAdminIfNotExists(dbmap)
}

func createAdminIfNotExists(dbmap *gorp.DbMap) {
	count, err := dbmap.SelectInt("select count(id) from users where username=$1", adminUsername)
	errors.CheckErr(err, "SelectInt failed")

	if count == 0 {
		admin := models.CreateUser(adminUsername, adminPassword)
		err = dbmap.Insert(&admin)
		errors.CheckErr(err, "Insert failed")
	}
}
