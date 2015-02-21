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
	dc := &DataContext{DBName: dbname, DBUser: dbuser}
	dc.OpenDb()
	defer dc.CloseDb()

	dropTablesIfExists(t, dc.DB)
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
	dc := &DataContext{DBName: dbname, DBUser: dbuser}
	dc.InitDb()
	defer dc.CloseDb()

	for _, table := range allTables {
		_, err := dc.DB.Exec(fmt.Sprintf(`SELECT * FROM %s LIMIT 1`, table))
		if err != nil {
			t.Error("Error querying table:", err)
		}
	}
}

func TestSeedDb(t *testing.T) {
	dc := &DataContext{DBName: dbname, DBUser: dbuser}
	dc.InitDb()
	defer dc.CloseDb()

	dc.SeedDb()

	var count int
	dc.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE username = $1`, adminUsername).Scan(&count)
	if count == 0 {
		t.Error("No admin user seeded")
	}

	dc.DB.QueryRow(`SELECT COUNT(*) FROM syntaxes`).Scan(&count)
	if count != len(models.LDAPv3Syntaxes) {
		t.Error("Wrong number of rows seeded")
	}

	dc.DB.QueryRow(`SELECT COUNT(*) FROM matching_rules`).Scan(&count)
	if count != len(models.LDAPv3MatchingRules) {
		t.Error("Wrong number of rows seeded")
	}

	dc.DB.QueryRow(`SELECT COUNT(*) FROM attribute_types`).Scan(&count)
	if count != len(models.LDAPv3AttributeTypes) {
		t.Error("Wrong number of rows seeded")
	}

	dc.DB.QueryRow(`SELECT COUNT(*) FROM object_classes`).Scan(&count)
	if count != len(models.LDAPv3ObjectClasses) {
		t.Error("Wrong number of rows seeded")
	}
}
