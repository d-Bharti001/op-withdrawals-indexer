package dbtypes

import (
	"database/sql/driver"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Bytes hexutil.Bytes

// Scan implements the database/sql Scanner interface
func (b *Bytes) Scan(value any) error {
	if value == nil {
		*b = nil
		return nil
	}

	v, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot scan %T into Bytes", value)
	}

	buf := make([]byte, len(v))
	copy(buf, v)
	*b = buf
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (b Bytes) Value() (driver.Value, error) {
	if b == nil {
		return nil, nil
	}
	return []byte(b), nil
}

func (b Bytes) Common() hexutil.Bytes {
	return hexutil.Bytes(b)
}

func (b Bytes) ByteSlice() []byte {
	if b == nil {
		return nil
	}
	return []byte(b)
}

type NullableBytes struct {
	Val *hexutil.Bytes
}

func NewNullableBytes(val *hexutil.Bytes) NullableBytes {
	return NullableBytes{Val: val}
}

// Scan implements the database/sql Scanner interface
func (b *NullableBytes) Scan(val any) error {
	if val == nil {
		b.Val = nil
		return nil
	}

	var bytes Bytes
	if err := bytes.Scan(val); err != nil {
		return err
	}

	commonBytes := bytes.Common()
	b.Val = &commonBytes
	return nil
}

// Value implements the database/sql/driver Valuer interface
func (b NullableBytes) Value() (driver.Value, error) {
	if b.Val == nil {
		return nil, nil
	}
	return []byte(*b.Val), nil
}

func (b NullableBytes) Common() *hexutil.Bytes {
	return b.Val
}

func (b NullableBytes) ByteSlice() []byte {
	if b.Val == nil {
		return nil
	}
	return []byte(*b.Val)
}
