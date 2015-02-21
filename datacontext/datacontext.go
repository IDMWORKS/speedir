package datacontext

import (
	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"

	"database/sql"

	// Imported for side-effects
	_ "github.com/lib/pq"
)

const (
	adminUsername = "admin"
	adminPassword = "admin"
)

// InitDb opens the DB & updates the schema as needed
func InitDb(dbname string, dbuser string) *sql.DB {
	db := OpenDb(dbname, dbuser)

	createTablesIfNotExists(db)

	return db
}

// OpenDb opens the database
func OpenDb(dbname string, dbuser string) *sql.DB {
	db, err := sql.Open("postgres", "user="+dbuser+" dbname="+dbname+" sslmode=disable")
	errors.CheckErr(err, "sql.Open failed")

	return db
}

func createTablesIfNotExists(db *sql.DB) {
	statements := []string{
		sqlCreateUsersTable,
	}
	for _, statement := range statements {
		_, err := db.Exec(statement)
		errors.CheckErr(err, "db.Exec failed")
	}
}

// SeedDb seeds the DB with data necessary for the app to run
func SeedDb(db *sql.DB) {
	createAdminIfNotExists(db)
}

func createAdminIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow(sqlSelectUserCountByUsername, adminUsername).Scan(&count)
	errors.CheckErr(err, "db.QueryRow failed")

	if count == 0 {
		admin := models.CreateUser(adminUsername, adminPassword)
		_, err := db.Exec(sqlInsertUserRow,
			admin.Created, admin.Username, admin.PasswordHash, admin.PasswordSalt)
		errors.CheckErr(err, "Insert failed")
	}
}
