package dbtypes

import (
	"database/sql/driver"
	"errors"
	"math/big"

	"github.com/jackc/pgtype"
)

var (
	u256Max = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	big10   = big.NewInt(10)

	ErrU256Overflow        = errors.New("number exceeds u256")
	ErrU256ContainsDecimal = errors.New("number contains fractional digits")
	ErrU256NotNull         = errors.New("number cannot be null")
)

//
// U256 (non-nullable NUMERIC)
//

type U256 struct {
	Int *big.Int
}

func NewU256(v *big.Int) U256 {
	return U256{Int: v}
}

// Scan implements the database/sql scanner interface
func (u *U256) Scan(src any) error {
	if src == nil {
		return errors.New("cannot scan null value as U256")
	}

	var numeric pgtype.Numeric
	if err := numeric.Scan(src); err != nil {
		return err
	}

	if numeric.Status == pgtype.Null {
		return ErrU256NotNull
	}
	if numeric.Exp < 0 {
		return ErrU256ContainsDecimal
	}

	num := numeric.Int
	if numeric.Exp > 0 {
		factor := new(big.Int).Exp(big10, big.NewInt(int64(numeric.Exp)), nil)
		num = new(big.Int).Mul(num, factor)
	}

	if num.Sign() < 0 || num.Cmp(u256Max) >= 0 {
		return ErrU256Overflow
	}

	u.Int = num
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (u U256) Value() (driver.Value, error) {
	if u.Int == nil {
		return nil, ErrU256NotNull
	}
	if u.Int.Sign() < 0 || u.Int.Cmp(u256Max) >= 0 {
		return nil, ErrU256Overflow
	}

	numeric := pgtype.Numeric{
		Int:    u.Int,
		Status: pgtype.Present,
	}
	return numeric.Value()
}

func (u U256) BigInt() *big.Int {
	return u.Int
}

//
// NullableU256 (nullable NUMERIC)
//

type NullableU256 struct {
	Int *big.Int
}

func NewNullableU256(v *big.Int) NullableU256 {
	return NullableU256{Int: v}
}

// Scan implements the database/sql scanner interface
func (n *NullableU256) Scan(src any) error {
	if src == nil {
		n.Int = nil
		return nil
	}

	var u U256
	if err := u.Scan(src); err != nil {
		return err
	}

	n.Int = u.Int
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (n NullableU256) Value() (driver.Value, error) {
	if n.Int == nil {
		return nil, nil
	}
	return U256{Int: n.Int}.Value()
}

func (n NullableU256) BigInt() *big.Int {
	return n.Int
}
