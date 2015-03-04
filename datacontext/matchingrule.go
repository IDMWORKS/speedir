package datacontext

import (
	"database/sql"
	"fmt"

	"github.com/idmworks/speedir/models"
)

// DBMatchingRule provides DB-centric methods for models.MatchingRule
type DBMatchingRule struct {
	*models.MatchingRule
}

// DBMatchingRulees provides DB-centric methods for a slice of DBMatchingRule
type DBMatchingRulees []*DBMatchingRule

// Scan scans each row from rows appending the populated matchingRule to matchingRules
func (matchingRules *DBMatchingRulees) scan(rows *sql.Rows) error {
	for rows.Next() {
		matchingRule := &DBMatchingRule{&models.MatchingRule{}}
		matchingRule.scan(rows)
		*matchingRules = append(*matchingRules, matchingRule)
	}
	return rows.Err()
}

// Scan scans the current row in rows to populate matchingRule
func (matchingRule *DBMatchingRule) scan(rows *sql.Rows) error {
	err := rows.Scan(
		&matchingRule.Name,
		&matchingRule.OID,
		&matchingRule.Syntax,
		&matchingRule.Names)
	return err
}

// SelectAllMatchingRulees returns a slice of DBMatchingRule for all matchingRules
func (dc *DataContext) SelectAllMatchingRules() (result DBMatchingRulees, err error) {
	matchingRules := make(DBMatchingRulees, 0)

	rows, err := dc.DB.Query(sqlSelectAllMatchingRules)
	if err != nil {
		return nil, fmt.Errorf("SelectAllMatchingRules failed: %v", err)
	}

	matchingRules.scan(rows)
	return matchingRules, nil
}
