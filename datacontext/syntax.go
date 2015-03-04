package datacontext

import (
	"database/sql"
	"fmt"

	"github.com/idmworks/speedir/models"
)

// DBSyntax provides DB-centric methods for models.Syntax
type DBSyntax struct {
	*models.Syntax
}

// DBSyntaxes provides DB-centric methods for a slice of DBSyntax
type DBSyntaxes []*DBSyntax

// Scan scans each row from rows appending the populated syntax to syntaxes
func (syntaxes *DBSyntaxes) scan(rows *sql.Rows) error {
	for rows.Next() {
		syntax := &DBSyntax{&models.Syntax{}}
		syntax.scan(rows)
		*syntaxes = append(*syntaxes, syntax)
	}
	return rows.Err()
}

// Scan scans the current row in rows to populate syntax
func (syntax *DBSyntax) scan(rows *sql.Rows) error {
	err := rows.Scan(
		&syntax.OID,
		&syntax.Description)
	return err
}

// SelectAllSyntaxes returns a slice of DBSyntax for all syntaxes
func (dc *DataContext) SelectAllSyntaxes() (result DBSyntaxes, err error) {
	syntaxes := make(DBSyntaxes, 0)

	rows, err := dc.DB.Query(sqlSelectAllSyntaxes)
	if err != nil {
		return nil, fmt.Errorf("SelectAllSyntaxes failed: %v", err)
	}

	syntaxes.scan(rows)
	return syntaxes, nil
}
