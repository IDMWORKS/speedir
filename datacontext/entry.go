package datacontext

import (
	"database/sql"
	"fmt"

	"github.com/idmworks/speedir/models"
)

// DBEntry provides DB-centric methods for models.Entry
type DBEntry struct {
	*models.Entry
}

// DBEntries provides DB-centric methods for a slice of DBEntry
type DBEntries []*DBEntry

// Scan scans each row from rows appending the populated entry to entries
func (entries *DBEntries) scan(rows *sql.Rows) error {
	for rows.Next() {
		entry := &DBEntry{&models.Entry{}}
		entry.scan(rows)
		*entries = append(*entries, entry)
	}
	return rows.Err()
}

// Scan scans the current row in rows to populate entry
func (entry *DBEntry) scan(rows *sql.Rows) error {
	err := rows.Scan(
		&entry.DN,
		&entry.Parent,
		&entry.RDN,
		&entry.Classes,
		&entry.UserValues,
		&entry.OperValues)
	return err
}

// SelectAllNamingContexts returns a slice of DBEntry for all Naming Contexts
func (dc *DataContext) SelectAllNamingContexts() (result DBEntries, err error) {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectAllNamingContexts)
	if err != nil {
		return nil, fmt.Errorf("SelectAllNamingContexts failed: %v", err)
	}

	entries.scan(rows)
	return entries, nil
}

// SelectEntriesByDN returns a slice of DBEntry with a matching DN
func (dc *DataContext) SelectEntriesByDN(dn string) (result DBEntries, err error) {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectEntriesByDN, dn)
	if err != nil {
		return nil, fmt.Errorf("SelectEntriesByDN failed: %v", err)
	}

	entries.scan(rows)
	return entries, nil
}

// SelectEntriesByParent returns a slice of DBEntry with a matching parent DN
func (dc *DataContext) SelectEntriesByParent(parent string) (result DBEntries, err error) {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectEntriesByParent, parent)
	if err != nil {
		return nil, fmt.Errorf("SelectEntriesByParent failed: %v", err)
	}

	entries.scan(rows)
	return entries, nil
}

// SelectEntriesByParent returns a slice of DBEntry with a matching parent DN
func (dc *DataContext) SelectEntryTreeByParent(parent string) (result DBEntries, err error) {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectEntryTreeByParent, parent+"%")
	if err != nil {
		return nil, fmt.Errorf("SelectEntryTreeByParent failed: %v", err)
	}

	entries.scan(rows)
	return entries, nil
}
