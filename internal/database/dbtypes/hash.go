package dbtypes

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Hash common.Hash

// Scan implements the database/sql scanner interface
func (h *Hash) Scan(value any) error {
	if value == nil {
		return errors.New("cannot scan null value as Hash")
	}

	var s string

	switch v := value.(type) {
	case string:
		s = strings.TrimSpace(v)
	case []byte:
		// Defensive: CHAR(66) may come as []byte
		s = strings.TrimSpace(string(v))
	default:
		return fmt.Errorf("cannot scan %T into Hash", value)
	}

	if !(len(s) == 66 && strings.HasPrefix(s, "0x")) ||
		(len(s) == 64 && !strings.HasPrefix(s, "0x")) {

		return fmt.Errorf("invalid hash string: %q", s)
	}

	*h = Hash(common.HexToHash(s))
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (h Hash) Value() (driver.Value, error) {
	hash := common.Hash(h)
	return hash.Hex(), nil
}

func (h Hash) Common() common.Hash {
	return common.Hash(h)
}

type NullableHash struct {
	Val *common.Hash
}

func NewNullableHash(val *common.Hash) NullableHash {
	return NullableHash{Val: val}
}

// Scan implements the database/sql Scanner interface
func (h *NullableHash) Scan(val any) error {
	if val == nil {
		h.Val = nil
		return nil
	}

	var hash Hash
	if err := hash.Scan(val); err != nil {
		return err
	}

	commonHash := hash.Common()
	h.Val = &commonHash
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (h NullableHash) Value() (driver.Value, error) {
	if h.Val == nil {
		return nil, nil
	}
	return h.Val.Hex(), nil
}

func (h NullableHash) Common() *common.Hash {
	return h.Val
}
