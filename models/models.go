package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
)

// StringSlice defines a slice of string for storage in a PG DB
type StringSlice []string

// Value converts a DB driver value (string) into a StringSlice
func (s StringSlice) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	buffer := bytes.NewBufferString("{")
	last := len(s) - 1
	for i, value := range s {
		buffer.WriteString(strconv.Quote(value))
		if i != last {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")

	return string(buffer.Bytes()), nil
}

// Scan scans a JSON encoded PG array as a StringSlice
// use array_to_json in SQL statements to return JSON
func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return error(errors.New("Scan value was not []bytes"))
	}

	var result StringSlice
	json.Unmarshal(bytes, &result)
	*s = result

	return nil
}

// AttributeValues defines a map of key-value pairs for storage in a PG DB
// TODO: does not handle mult-value attributes
type AttributeValues map[string][]string

// Value converts a DB driver value (string) into a StringSlice
func (s AttributeValues) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

// Scan scans a JSON(B) value from PG array as AttributeValues
func (s *AttributeValues) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return error(errors.New("Scan value was not []bytes"))
	}

	var result AttributeValues
	json.Unmarshal(bytes, &result)
	*s = result

	return nil
}
