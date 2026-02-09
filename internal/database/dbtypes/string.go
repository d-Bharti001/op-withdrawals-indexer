package dbtypes

import (
	"database/sql/driver"
	"fmt"
)

type NullableString struct {
	Val *string
}

func NewNullableString(s *string) NullableString {
	return NullableString{Val: s}
}

// Scan implements the database/sql scanner interface
func (n *NullableString) Scan(src any) error {
	if src == nil {
		n.Val = nil
		return nil
	}

	var str string

	switch s := src.(type) {
	case []byte:
		str = string(s)

	case string:
		str = s

	default:
		return fmt.Errorf("cannot scan %T into NullableString", src)
	}

	n.Val = &str
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (n NullableString) Value() (driver.Value, error) {
	if n.Val == nil {
		return nil, nil
	}
	return *n.Val, nil
}

func (n NullableString) IsValid() bool {
	if n.Val == nil {
		return false
	}
	return true
}
