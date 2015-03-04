package datacontext

import (
	"database/sql"
	"fmt"

	"github.com/idmworks/speedir/models"
)

// DBObjectClass provides DB-centric methods for models.ObjectClass
type DBObjectClass struct {
	*models.ObjectClass
}

// DBObjectClasses provides DB-centric methods for a slice of DBObjectClass
type DBObjectClasses []*DBObjectClass

// Scan scans each row from rows appending the populated objectClass to objectClasses
func (objectClasses *DBObjectClasses) scan(rows *sql.Rows) error {
	for rows.Next() {
		objectClass := &DBObjectClass{&models.ObjectClass{}}
		objectClass.scan(rows)
		*objectClasses = append(*objectClasses, objectClass)
	}
	return rows.Err()
}

// Scan scans the current row in rows to populate objectClass
func (objectClass *DBObjectClass) scan(rows *sql.Rows) error {
	err := rows.Scan(
		&objectClass.Name,
		&objectClass.OID,
		&objectClass.Super,
		&objectClass.Names,
		&objectClass.Flags,
		&objectClass.MustAttributes,
		&objectClass.MayAttributes)
	if err != nil {
		return err
	}

	if !objectClass.Super.Valid {
		objectClass.Super = sql.NullString{String: models.TopClass, Valid: true}
	}
	return nil
}

// SelectAllObjectClasses returns a slice of DBObjectClass for all objectClasses
func (dc *DataContext) SelectAllObjectClasses() (result DBObjectClasses, err error) {
	objectClasses := make(DBObjectClasses, 0)

	rows, err := dc.DB.Query(sqlSelectAllObjectClasses)
	if err != nil {
		return nil, fmt.Errorf("SelectAllObjectClasses failed: %v", err)
	}

	objectClasses.scan(rows)
	return objectClasses, nil
}
