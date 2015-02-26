package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

// DBAttributeType provides DB-centric methods for models.AttributeType
type DBAttributeType struct {
	*models.AttributeType
}

// DBAttributeTypees provides DB-centric methods for a slice of DBAttributeType
type DBAttributeTypees []*DBAttributeType

// Scan scans each row from rows appending the populated attributeType to attributeTypes
func (attributeTypes *DBAttributeTypees) scan(rows *sql.Rows) {
	for rows.Next() {
		attributeType := &DBAttributeType{&models.AttributeType{}}
		attributeType.scan(rows)
		*attributeTypes = append(*attributeTypes, attributeType)
	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

// Scan scans the current row in rows to populate attributeType
func (attributeType *DBAttributeType) scan(rows *sql.Rows) {
	err := rows.Scan(
		&attributeType.Name,
		&attributeType.OID,
		&attributeType.Super,
		&attributeType.Syntax,
		&attributeType.Names,
		&attributeType.Flags,
		&attributeType.Usage,
		&attributeType.EqualityMatch,
		&attributeType.SubstrMatch,
		&attributeType.OrderingMatch)
	errors.CheckErr(err, "rows.Scan failed")
}

// SelectAllAttributeTypees returns a slice of DBAttributeType for all attributeTypes
func (dc *DataContext) SelectAllAttributeTypes() DBAttributeTypees {
	attributeTypes := make(DBAttributeTypees, 0)

	rows, err := dc.DB.Query(sqlSelectAllAttributeTypes)
	errors.CheckErr(err, "Select failed")
	attributeTypes.scan(rows)

	return attributeTypes
}
