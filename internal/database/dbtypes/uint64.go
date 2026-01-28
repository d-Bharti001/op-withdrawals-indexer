package dbtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"strconv"
)

var ErrUint64Negative = errors.New("negative value for uint64")

type NullableUint64 struct {
	Val *uint64
}

func NewNullableUint64(v *uint64) NullableUint64 {
	return NullableUint64{Val: v}
}

// Scan implements the database/sql scanner interface
func (n *NullableUint64) Scan(src any) error {
	if src == nil {
		n.Val = nil
		return nil
	}

	var v int64

	switch x := src.(type) {
	case int64:
		v = x

	case int32:
		v = int64(x)

	case []byte:
		parsed, err := strconv.ParseInt(string(x), 10, 64)
		if err != nil {
			return err
		}
		v = parsed

	case string:
		parsed, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return err
		}
		v = parsed

	default:
		return fmt.Errorf("cannot scan %T into NullableUint64", src)
	}

	if v < 0 {
		return ErrUint64Negative
	}

	u := uint64(v)
	n.Val = &u
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (n NullableUint64) Value() (driver.Value, error) {
	if n.Val == nil {
		return nil, nil
	}
	if *n.Val > math.MaxInt64 {
		return nil, fmt.Errorf("uint64 value %d overflows int64", *n.Val)
	}
	return int64(*n.Val), nil
}
