package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

// DBMatchingRule provides DB-centric methods for models.MatchingRule
type DBMatchingRule struct {
	*models.MatchingRule
}

// DBMatchingRulees provides DB-centric methods for a slice of DBMatchingRule
type DBMatchingRulees []*DBMatchingRule

// Scan scans each row from rows appending the populated matchingRule to matchingRules
func (matchingRules *DBMatchingRulees) scan(rows *sql.Rows) {
	for rows.Next() {
		matchingRule := &DBMatchingRule{&models.MatchingRule{}}
		matchingRule.scan(rows)
		*matchingRules = append(*matchingRules, matchingRule)
	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

// Scan scans the current row in rows to populate matchingRule
func (matchingRule *DBMatchingRule) scan(rows *sql.Rows) {
	err := rows.Scan(
		&matchingRule.Name,
		&matchingRule.OID,
		&matchingRule.Syntax,
		&matchingRule.Names)
	errors.CheckErr(err, "rows.Scan failed")
}

// SelectAllMatchingRulees returns a slice of DBMatchingRule for all matchingRules
func (dc *DataContext) SelectAllMatchingRules() DBMatchingRulees {
	matchingRules := make(DBMatchingRulees, 0)

	rows, err := dc.DB.Query(sqlSelectAllMatchingRules)
	errors.CheckErr(err, "Select failed")
	matchingRules.scan(rows)

	return matchingRules
}
