package models

import (
	"bytes"
	"database/sql/driver"
	"strconv"
)

type StringSlice []string

func (names StringSlice) Value() (driver.Value, error) {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	last := len(names) - 1
	for i, name := range names {
		buffer.WriteString(strconv.Quote(name))
		if i != last {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")

	return string(buffer.Bytes()), nil
}
