package dbtypes

import (
	"database/sql/driver"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Bytes hexutil.Bytes

// Scan implements the database/sql scanner interface
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

func (b Bytes) HexUtilBytes() hexutil.Bytes {
	return hexutil.Bytes(b)
}

func (b Bytes) ByteSlice() []byte {
	if b == nil {
		return nil
	}
	return []byte(b)
}
