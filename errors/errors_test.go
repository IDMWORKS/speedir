package errors

import (
	"testing"
)

func TestCheckErr(t *testing.T) {
	CheckErr(nil, "")
}
