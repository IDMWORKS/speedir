package datacontext

import (
	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"

	"database/sql"

	// Imported for side-effects
	"fmt"
	_ "github.com/lib/pq"
)

const (
	adminUsername = "admin"
	adminPassword = "admin"
)

type DataContext struct {
	DBName string
	DBUser string
	DB     *sql.DB
}

// InitDb opens the DB & updates the schema as needed
func (dc *DataContext) InitDb() {
	dc.OpenDb()

	createTablesIfNotExists(dc.DB)
}

// OpenDb opens the database
func (dc *DataContext) OpenDb() {
	db, err := sql.Open("postgres", "user="+dc.DBUser+" dbname="+dc.DBName+" sslmode=disable")
	errors.CheckErr(err, "sql.Open failed")
	dc.DB = db
}

// OpenDb opens the database
func (dc *DataContext) CloseDb() {
	dc.DB.Close()
	dc.DB = nil
}

// SeedDb seeds the DB with data necessary for the app to run
func (dc *DataContext) SeedDb() {
	createAdminIfNotExists(dc.DB)
	createSyntaxesIfNotExists(dc.DB)
	createMatchingRulesIfNotExists(dc.DB)
	createAttributeTypesIfNotExists(dc.DB)
	createObjectClassesIfNotExists(dc.DB)
	createDummySchemaIfNotExists(dc.DB)
}

func createTablesIfNotExists(db *sql.DB) {
	statements := []string{
		sqlCreateUsersTable,
		sqlCreateSyntaxesTable,
		sqlCreateMatchingRulesTable,
		sqlCreateAttributeTypesTable,
		sqlCreateObjectClassesTable,
		sqlCreateEntriesTable,
	}
	for _, statement := range statements {
		_, err := db.Exec(statement)
		errors.CheckErr(err, "db.Exec failed")
	}
}

func createDummySchemaIfNotExists(db *sql.DB) {
	var count int
	err := db.QueryRow(sqlSelectEntryCount).Scan(&count)
	errors.CheckErr(err, "db.QueryRow failed")

	if count == 0 {
		rootDN := "dc=example,dc=org"

		insertEntryRow(db, &models.Entry{
			DN:      rootDN,
			Parent:  sql.NullString{Valid: false},
			RDN:     rootDN,
			Classes: models.StringSlice{models.DomainClass},
			UserValues: models.AttributeValues{
				models.DomainComponentAttribute: []string{"example"},
				models.CommonNameAttribute:      []string{"example"},
			},
		})

		parentDN := rootDN
		commonName := "Users"
		commonNameEntry := fmt.Sprintf("%s=%s", models.CommonNameAttribute, commonName)
		entryDN := commonNameEntry + "," + parentDN

		insertEntryRow(db, &models.Entry{
			DN:      entryDN,
			Parent:  sql.NullString{String: parentDN, Valid: true},
			RDN:     commonNameEntry,
			Classes: models.StringSlice{models.GroupOfNamesClass},
			UserValues: models.AttributeValues{
				models.CommonNameAttribute: []string{commonName},
			},
		})

		groupDN := entryDN

		parentDN = groupDN
		commonName = "Test User"
		commonNameEntry = fmt.Sprintf("%s=%s", models.CommonNameAttribute, commonName)
		entryDN = commonNameEntry + "," + parentDN

		insertEntryRow(db, &models.Entry{
			DN:      entryDN,
			Parent:  sql.NullString{String: parentDN, Valid: true},
			RDN:     commonNameEntry,
			Classes: models.StringSlice{models.PersonClass},
			UserValues: models.AttributeValues{
				models.CommonNameAttribute: []string{commonName},
				models.SurnameAttribute:    []string{commonName},
			},
		})

		parentDN = groupDN
		commonName = "Test User2"
		commonNameEntry = fmt.Sprintf("%s=%s", models.CommonNameAttribute, commonName)
		entryDN = commonNameEntry + "," + parentDN

		insertEntryRow(db, &models.Entry{
			DN:      entryDN,
			Parent:  sql.NullString{String: parentDN, Valid: true},
			RDN:     commonNameEntry,
			Classes: models.StringSlice{models.PersonClass},
			UserValues: models.AttributeValues{
				models.CommonNameAttribute: []string{commonName},
				models.SurnameAttribute:    []string{commonName},
			},
		})
	}
}

func insertEntryRow(db *sql.DB, entry *models.Entry) {
	_, err := db.Exec(
		sqlInsertEntryRow,
		entry.DN,
		entry.Parent,
		entry.RDN,
		entry.Classes,
		entry.UserValues,
		entry.OperValues)
	errors.CheckErr(err, "Insert failed")
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
				rule.Name, rule.OID, rule.Syntax, rule.Names)
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
				attr.Name, attr.OID, attr.Syntax, attr.Super, attr.Names, attr.Flags,
				attr.Usage, attr.EqualityMatch, attr.SubstrMatch, attr.OrderingMatch)
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
				class.Name, class.OID, class.Super, class.Names, class.Flags,
				class.MustAttributes, class.MayAttributes)
			errors.CheckErr(err, "Insert failed")
		}
	}
}
