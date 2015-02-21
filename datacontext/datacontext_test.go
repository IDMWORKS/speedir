package datacontext

import (
	"fmt"
	"testing"

	"github.com/idmworks/speedir/models"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	dbname = "speedir_test"
	dbuser = "speedir_test"
)

var allTables = []string{
	"users",
}

func TestMain(t *testing.T) {
	db := OpenDb(dbname, dbuser)
	defer db.Close()

	dropTablesIfExists(t, db)
}

func dropTablesIfExists(t *testing.T, db *sql.DB) {
	for _, table := range allTables {
		_, err := db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table))
		if err != nil {
			t.Error("Error querying users table:", err)
		}
	}
}

func TestInitDb(t *testing.T) {
	db := InitDb(dbname, dbuser)
	defer db.Close()

	for _, table := range allTables {
		_, err := db.Exec(fmt.Sprintf(`SELECT * FROM %s LIMIT 1`, table))
		if err != nil {
			t.Error("Error querying table:", err)
		}
	}
}

func TestSeedDb(t *testing.T) {
	db := InitDb(dbname, dbuser)
	defer db.Close()

	SeedDb(db)

	var count int
	db.QueryRow(`SELECT COUNT(*) FROM users WHERE username = $1`, adminUsername).Scan(&count)
	if count == 0 {
		t.Error("No admin user seeded")
	}
}

func openPgDb(dbname string, dbuser string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbuser, dbname))
	errors.CheckErr(err, "sql.Open failed")
	return db
}
