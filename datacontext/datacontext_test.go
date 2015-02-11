package datacontext

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/idmworks/speedir/errors"
	"gopkg.in/gorp.v1"
)

const (
	dbname = "speedir_test"
	dbuser = "speedir_test"
)

func TestMain(t *testing.T) {
	db := openPgDb(dbname, dbuser)

	// drop existing mapped tables before tests run
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.DropTablesIfExists()
}

func TestInitDb(t *testing.T) {
	InitDb(dbname, dbuser)

	db := openPgDb(dbname, dbuser)
	_, err := db.Query("SELECT * FROM users LIMIT 1;")
	if err != nil {
		t.Error("Error querying users table:", err)
	}
}

func TestSeedDb(t *testing.T) {
	dbmap := InitDb(dbname, dbuser)
	SeedDb(dbmap)

	db := openPgDb(dbname, dbuser)
	rows, err := db.Query("SELECT * FROM users WHERE username = $1;", "admin")
	if err != nil {
		t.Error("Error querying users table:", err)
	}
	if !rows.Next() {
		t.Error("No admin user seeded")
	}
}

func openPgDb(dbname string, dbuser string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbuser, dbname))
	errors.CheckErr(err, "sql.Open failed")
	return db
}
