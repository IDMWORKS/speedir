package datacontext

import (
	"database/sql"
	"fmt"
	// Imported for side-effects
	_ "github.com/lib/pq"

	"github.com/idmworks/speedir/models"
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
func (dc *DataContext) InitDb() error {
	if err := dc.OpenDb(); err != nil {
		return err
	}

	return createTablesIfNotExists(dc.DB)
}

// OpenDb opens the database
func (dc *DataContext) OpenDb() error {
	db, err := sql.Open("postgres", "user="+dc.DBUser+" dbname="+dc.DBName+" sslmode=disable")
	dc.DB = db
	return err
}

// OpenDb opens the database
func (dc *DataContext) CloseDb() {
	dc.DB.Close()
	dc.DB = nil
}

// SeedDb seeds the DB with data necessary for the app to run
func (dc *DataContext) SeedDb() error {
	if err := createAdminIfNotExists(dc.DB); err != nil {
		return err
	}
	if err := createSyntaxesIfNotExists(dc.DB); err != nil {
		return err
	}
	if err := createMatchingRulesIfNotExists(dc.DB); err != nil {
		return err
	}
	if err := createAttributeTypesIfNotExists(dc.DB); err != nil {
		return err
	}
	if err := createObjectClassesIfNotExists(dc.DB); err != nil {
		return err
	}
	if err := createDummySchemaIfNotExists(dc.DB); err != nil {
		return err
	}
	return nil
}

func createTablesIfNotExists(db *sql.DB) error {
	statements := []string{
		sqlCreateUsersTable,
		sqlCreateSyntaxesTable,
		sqlCreateMatchingRulesTable,
		sqlCreateAttributeTypesTable,
		sqlCreateObjectClassesTable,
		sqlCreateEntriesTable,
	}
	for _, statement := range statements {
		if _, err := db.Exec(statement); err != nil {
			return err
		}
	}
	return nil
}

func createDummySchemaIfNotExists(db *sql.DB) error {
	var count int
	if err := db.QueryRow(sqlSelectEntryCount).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		rootDN := "dc=example,dc=org"

		if err := insertEntryRow(db, &models.Entry{
			DN:      rootDN,
			Parent:  sql.NullString{Valid: false},
			RDN:     rootDN,
			Classes: models.StringSlice{models.DomainClass},
			UserValues: models.AttributeValues{
				models.DomainComponentAttribute: []string{"example"},
			},
		}); err != nil {
			return err
		}

		parentDN := rootDN
		commonName := "Users"
		commonNameEntry := fmt.Sprintf("%s=%s", models.CommonNameAttribute, commonName)
		entryDN := commonNameEntry + "," + parentDN

		if err := insertEntryRow(db, &models.Entry{
			DN:      entryDN,
			Parent:  sql.NullString{String: parentDN, Valid: true},
			RDN:     commonNameEntry,
			Classes: models.StringSlice{models.GroupOfNamesClass},
			UserValues: models.AttributeValues{
				models.CommonNameAttribute: []string{commonName},
			},
		}); err != nil {
			return err
		}

		groupDN := entryDN

		parentDN = groupDN
		commonName = "Test User"
		commonNameEntry = fmt.Sprintf("%s=%s", models.CommonNameAttribute, commonName)
		entryDN = commonNameEntry + "," + parentDN

		if err := insertEntryRow(db, &models.Entry{
			DN:      entryDN,
			Parent:  sql.NullString{String: parentDN, Valid: true},
			RDN:     commonNameEntry,
			Classes: models.StringSlice{models.PersonClass},
			UserValues: models.AttributeValues{
				models.CommonNameAttribute: []string{commonName},
				models.SurnameAttribute:    []string{commonName},
			},
		}); err != nil {
			return err
		}

		parentDN = groupDN
		commonName = "Test User2"
		commonNameEntry = fmt.Sprintf("%s=%s", models.CommonNameAttribute, commonName)
		entryDN = commonNameEntry + "," + parentDN

		if err := insertEntryRow(db, &models.Entry{
			DN:      entryDN,
			Parent:  sql.NullString{String: parentDN, Valid: true},
			RDN:     commonNameEntry,
			Classes: models.StringSlice{models.PersonClass},
			UserValues: models.AttributeValues{
				models.CommonNameAttribute: []string{commonName},
				models.SurnameAttribute:    []string{commonName},
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

func insertEntryRow(db *sql.DB, entry *models.Entry) error {
	_, err := db.Exec(
		sqlInsertEntryRow,
		entry.DN,
		entry.Parent,
		entry.RDN,
		entry.Classes,
		entry.UserValues,
		entry.OperValues)
	return err
}

func createAdminIfNotExists(db *sql.DB) error {
	var count int
	if err := db.QueryRow(sqlSelectUserCountByUsername, adminUsername).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		admin := models.CreateUser(adminUsername, adminPassword)
		if _, err := db.Exec(sqlInsertUserRow,
			admin.Created, admin.Username, admin.PasswordHash, admin.PasswordSalt); err != nil {
			return err
		}
	}

	return nil
}

func createSyntaxesIfNotExists(db *sql.DB) error {
	var count int
	if err := db.QueryRow(sqlSelectSyntaxCount).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		for _, syntax := range models.LDAPv3Syntaxes {
			if _, err := db.Exec(sqlInsertSyntaxRow,
				syntax.OID, syntax.Description); err != nil {
				return err
			}
		}
	}

	return nil
}

func createMatchingRulesIfNotExists(db *sql.DB) error {
	var count int
	if err := db.QueryRow(sqlSelectMatchingRuleCount).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		for _, rule := range models.LDAPv3MatchingRules {
			if _, err := db.Exec(sqlInsertMatchingRuleRow,
				rule.Name, rule.OID, rule.Syntax, rule.Names); err != nil {
				return err
			}
		}
	}

	return nil
}

func createAttributeTypesIfNotExists(db *sql.DB) error {
	var count int
	if err := db.QueryRow(sqlSelectAttributeTypeCount).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		for _, attr := range models.LDAPv3AttributeTypes {
			if _, err := db.Exec(sqlInsertAttributeTypeRow,
				attr.Name, attr.OID, attr.Syntax, attr.Super, attr.Names, attr.Flags,
				attr.Usage, attr.EqualityMatch, attr.SubstrMatch, attr.OrderingMatch); err != nil {
				return err
			}
		}
	}

	return nil
}

func createObjectClassesIfNotExists(db *sql.DB) error {
	var count int
	if err := db.QueryRow(sqlSelectObjectClassCount).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		for _, class := range models.LDAPv3ObjectClasses {
			if _, err := db.Exec(sqlInsertObjectClassRow,
				class.Name, class.OID, class.Super, class.Names, class.Flags,
				class.MustAttributes, class.MayAttributes); err != nil {
				return err
			}
		}
	}

	return nil
}
