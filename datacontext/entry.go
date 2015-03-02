package datacontext

import (
	"database/sql"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
)

// DBEntry provides DB-centric methods for models.Entry
type DBEntry struct {
	*models.Entry
}

// DBEntries provides DB-centric methods for a slice of DBEntry
type DBEntries []*DBEntry

// Scan scans each row from rows appending the populated entry to entries
func (entries *DBEntries) scan(rows *sql.Rows) {
	for rows.Next() {
		entry := &DBEntry{&models.Entry{}}
		entry.scan(rows)
		*entries = append(*entries, entry)
	}
	errors.CheckErr(rows.Err(), "rows.Next failed")
}

// Scan scans the current row in rows to populate entry
func (entry *DBEntry) scan(rows *sql.Rows) {
	err := rows.Scan(
		&entry.DN,
		&entry.Parent,
		&entry.RDN,
		&entry.Classes,
		&entry.UserValues,
		&entry.OperValues)
	errors.CheckErr(err, "rows.Scan failed")
}

// SelectAllNamingContexts returns a slice of DBEntry for all Naming Contexts
func (dc *DataContext) SelectAllNamingContexts() DBEntries {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectAllNamingContexts)
	errors.CheckErr(err, "Select failed")
	entries.scan(rows)

	return entries
}

// SelectEntriesByDN returns a slice of DBEntry with a matching DN
func (dc *DataContext) SelectEntriesByDN(dn string) DBEntries {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectEntriesByDN, dn)
	errors.CheckErr(err, "Select failed")
	entries.scan(rows)

	return entries
}

// SelectEntriesByParent returns a slice of DBEntry with a matching parent DN
func (dc *DataContext) SelectEntriesByParent(parent string) DBEntries {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectEntriesByParent, parent)
	errors.CheckErr(err, "Select failed")
	entries.scan(rows)

	return entries
}

// SelectEntriesByParent returns a slice of DBEntry with a matching parent DN
func (dc *DataContext) SelectEntryTreeByParent(parent string) DBEntries {
	entries := make(DBEntries, 0)

	rows, err := dc.DB.Query(sqlSelectEntryTreeByParent, parent+"%")
	errors.CheckErr(err, "Select failed")
	entries.scan(rows)

	return entries
}
