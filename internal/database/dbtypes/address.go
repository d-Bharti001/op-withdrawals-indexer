package dbtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address

// Scan implements the database/sql Scanner interface
func (a *Address) Scan(value any) error {
	if value == nil {
		return errors.New("cannot scan null value as Address")
	}

	switch v := value.(type) {
	case string:
		s := strings.TrimSpace(v)
		if !common.IsHexAddress(s) {
			return fmt.Errorf("invalid ethereum address: %q", s)
		}
		*a = Address(common.HexToAddress(s))
		return nil

	case []byte:
		s := strings.TrimSpace(string(v))
		if !common.IsHexAddress(s) {
			return fmt.Errorf("invalid ethereum address: %q", s)
		}
		*a = Address(common.HexToAddress(s))
		return nil

	default:
		return fmt.Errorf("cannot scan %T into Address", value)
	}
}

// Value implements the database/sql/driver Valuer interface
func (a Address) Value() (driver.Value, error) {
	addr := common.Address(a)
	return strings.ToLower(addr.Hex()), nil
}

func (a Address) Common() common.Address {
	return common.Address(a)
}

type NullableAddress struct {
	Val *common.Address
}

func NewNullableAddress(val *common.Address) NullableAddress {
	return NullableAddress{Val: val}
}

// Scan implements the database/sql Scanner interface
func (a *NullableAddress) Scan(val any) error {
	if val == nil {
		a.Val = nil
		return nil
	}

	var addr Address
	if err := addr.Scan(val); err != nil {
		return err
	}

	commonAddr := common.Address(addr)
	a.Val = &commonAddr
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (a *NullableAddress) Value() (driver.Value, error) {
	if a.Val == nil {
		return nil, nil
	}
	return strings.ToLower(a.Val.Hex()), nil
}

func (a *NullableAddress) Common() *common.Address {
	return a.Val
}
