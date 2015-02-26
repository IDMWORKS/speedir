package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

// DBSyntax provides DB-centric methods for models.Syntax
type DBSyntax struct {
	*models.Syntax
}

// DBSyntaxes provides DB-centric methods for a slice of DBSyntax
type DBSyntaxes []*DBSyntax

// Scan scans each row from rows appending the populated syntax to syntaxes
func (syntaxes *DBSyntaxes) scan(rows *sql.Rows) {
	for rows.Next() {
		syntax := &DBSyntax{&models.Syntax{}}
		syntax.scan(rows)
		*syntaxes = append(*syntaxes, syntax)
	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

// Scan scans the current row in rows to populate syntax
func (syntax *DBSyntax) scan(rows *sql.Rows) {
	err := rows.Scan(
		&syntax.OID,
		&syntax.Description)
	errors.CheckErr(err, "rows.Scan failed")
}

// SelectAllSyntaxes returns a slice of DBSyntax for all syntaxes
func (dc *DataContext) SelectAllSyntaxes() DBSyntaxes {
	syntaxes := make(DBSyntaxes, 0)

	rows, err := dc.DB.Query(sqlSelectAllSyntaxes)
	errors.CheckErr(err, "Select failed")
	syntaxes.scan(rows)

	return syntaxes
}
