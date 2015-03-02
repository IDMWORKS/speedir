package models

import (
	"database/sql"
)

// Entry model in the DB
type Entry struct {
	DN         string
	Parent     sql.NullString
	RDN        string
	Classes    StringSlice
	UserValues AttributeValues
	OperValues AttributeValues
}
