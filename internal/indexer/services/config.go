package indexer

import (
	"time"
)

var DefaultL1PollingInterval = 5 * time.Second
var DefaultL2PollingInterval = 5 * time.Second

const ShutdownTimeout = 10 * time.Second

type IndexerInitConfig struct {
	L1RPCUrl                 string
	L1ChainID                uint64
	L1ChainName              string
	L1BlockConfirmationDepth uint64

	L2RPCUrl                 string
	L2ChainID                uint64
	L2ChainName              string
	L2BlockConfirmationDepth uint64

	SystemConfigAddr string

	DBConnStr     string
	DBMaxOpenConn int
	DBMaxIdleConn int
	DBMaxIdleTime string
}
