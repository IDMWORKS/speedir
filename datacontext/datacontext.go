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
		sqlCreateSyntaxesTable,
		sqlCreateMatchingRulesTable,
		sqlCreateAttributeTypesTable,
		sqlCreateObjectClassesTable,
	}
	for _, statement := range statements {
		_, err := db.Exec(statement)
		errors.CheckErr(err, "db.Exec failed")
	}
}

// SeedDb seeds the DB with data necessary for the app to run
func SeedDb(db *sql.DB) {
	createAdminIfNotExists(db)
	createSyntaxesIfNotExists(db)
	createMatchingRulesIfNotExists(db)
	createAttributeTypesIfNotExists(db)
	createObjectClassesIfNotExists(db)
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

func createSyntaxesIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow(sqlSelectSyntaxCount).Scan(&count)
	errors.CheckErr(err, "db.QueryRow failed")

	if count == 0 {
		for _, syntax := range models.LDAPv3Syntaxes {
			_, err := db.Exec(sqlInsertSyntaxRow,
				syntax.OID, syntax.Description)
			errors.CheckErr(err, "Insert failed")
		}
	}
}

func createMatchingRulesIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow(sqlSelectMatchingRuleCount).Scan(&count)
	errors.CheckErr(err, "db.QueryRow failed")

	if count == 0 {
		for _, rule := range models.LDAPv3MatchingRules {
			_, err := db.Exec(sqlInsertMatchingRuleRow,
				rule.OID, rule.SyntaxID, rule.Names)
			errors.CheckErr(err, "Insert failed")
		}
	}
}

func createAttributeTypesIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow(sqlSelectAttributeTypeCount).Scan(&count)
	errors.CheckErr(err, "db.QueryRow failed")

	if count == 0 {
		for _, attr := range models.LDAPv3AttributeTypes {
			_, err := db.Exec(sqlInsertAttributeTypeRow,
				attr.OID, attr.SyntaxID, attr.SuperID, attr.Names, attr.Flags,
				attr.EqualityMatchID, attr.SubstrMatchID, attr.OrderingMatchID)
			errors.CheckErr(err, "Insert failed")
		}
	}
}

func createObjectClassesIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow(sqlSelectObjectClassCount).Scan(&count)
	errors.CheckErr(err, "db.QueryRow failed")

	if count == 0 {
		for _, class := range models.LDAPv3ObjectClasses {
			_, err := db.Exec(sqlInsertObjectClassRow,
				class.OID, class.SuperID, class.Names, class.Flags,
				class.MustAttributes, class.MayAttributes)
			errors.CheckErr(err, "Insert failed")
		}
	}
}
