package datacontext

import (
	"database/sql"
	"fmt"

	"github.com/idmworks/speedir/models"
)

// DBAttributeType provides DB-centric methods for models.AttributeType
type DBAttributeType struct {
	*models.AttributeType
}

// DBAttributeTypees provides DB-centric methods for a slice of DBAttributeType
type DBAttributeTypees []*DBAttributeType

// Scan scans each row from rows appending the populated attributeType to attributeTypes
func (attributeTypes *DBAttributeTypees) scan(rows *sql.Rows) error {
	for rows.Next() {
		attributeType := &DBAttributeType{&models.AttributeType{}}
		attributeType.scan(rows)
		*attributeTypes = append(*attributeTypes, attributeType)
	}
	return rows.Err()
}

// Scan scans the current row in rows to populate attributeType
func (attributeType *DBAttributeType) scan(rows *sql.Rows) error {
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
	return err
}

// SelectAllAttributeTypees returns a slice of DBAttributeType for all attributeTypes
func (dc *DataContext) SelectAllAttributeTypes() (result DBAttributeTypees, err error) {
	attributeTypes := make(DBAttributeTypees, 0)

	rows, err := dc.DB.Query(sqlSelectAllAttributeTypes)
	if err != nil {
		return nil, fmt.Errorf("SelectAllAttributeTypes failed: %v", err)
	}

	attributeTypes.scan(rows)
	return attributeTypes, nil
}
