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
	"object_classes",
	"attribute_types",
	"matching_rules",
	"syntaxes",
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

	db.QueryRow(`SELECT COUNT(*) FROM syntaxes`).Scan(&count)
	if count != len(models.LDAPv3Syntaxes) {
		t.Error("Wrong number of rows seeded")
	}

	db.QueryRow(`SELECT COUNT(*) FROM matching_rules`).Scan(&count)
	if count != len(models.LDAPv3MatchingRules) {
		t.Error("Wrong number of rows seeded")
	}

	db.QueryRow(`SELECT COUNT(*) FROM attribute_types`).Scan(&count)
	if count != len(models.LDAPv3AttributeTypes) {
		t.Error("Wrong number of rows seeded")
	}

	db.QueryRow(`SELECT COUNT(*) FROM object_classes`).Scan(&count)
	if count != len(models.LDAPv3ObjectClasses) {
		t.Error("Wrong number of rows seeded")
	}
}
