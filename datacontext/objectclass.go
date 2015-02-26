package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

// DBObjectClass provides DB-centric methods for models.ObjectClass
type DBObjectClass struct {
	*models.ObjectClass
}

// DBObjectClasses provides DB-centric methods for a slice of DBObjectClass
type DBObjectClasses []*DBObjectClass

// Scan scans each row from rows appending the populated objectClass to objectClasses
func (objectClasses *DBObjectClasses) scan(rows *sql.Rows) {
	for rows.Next() {
		objectClass := &DBObjectClass{&models.ObjectClass{}}
		objectClass.scan(rows)
		*objectClasses = append(*objectClasses, objectClass)
	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

// Scan scans the current row in rows to populate objectClass
func (objectClass *DBObjectClass) scan(rows *sql.Rows) {
	err := rows.Scan(
		&objectClass.Name,
		&objectClass.OID,
		&objectClass.Super,
		&objectClass.Names,
		&objectClass.Flags,
		&objectClass.MustAttributes,
		&objectClass.MayAttributes)
	errors.CheckErr(err, "rows.Scan failed")
	if !objectClass.Super.Valid {
		objectClass.Super = sql.NullString{String: models.TopClass, Valid: true}
	}
}

// SelectAllObjectClasses returns a slice of DBObjectClass for all objectClasses
func (dc *DataContext) SelectAllObjectClasses() DBObjectClasses {
	objectClasses := make(DBObjectClasses, 0)

	rows, err := dc.DB.Query(sqlSelectAllObjectClasses)
	errors.CheckErr(err, "Select failed")
	objectClasses.scan(rows)

	return objectClasses
}
