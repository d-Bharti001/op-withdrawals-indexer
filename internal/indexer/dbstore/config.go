package dbstore

import (
	"errors"
	"time"
)

const DBQueryTimeout = 3 * time.Second

var ErrNotFound = errors.New("resource not found")
