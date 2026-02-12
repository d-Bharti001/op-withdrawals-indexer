package indexer

import (
	"time"
)

const DefaultL1ChainUnstableBlocksDepth uint64 = 24
const DefaultL2ChainUnstableBlocksDepth uint64 = 20

const DefaultL1BlockScanBatchSizeLimit uint64 = 250
const DefaultL2BlockScanBatchSizeLimit uint64 = 1000

const DefaultL1PollingInterval = 5 * time.Second
const DefaultL2PollingInterval = 2 * time.Second

const ShutdownTimeout = 10 * time.Second

type IndexerInitConfig struct {
	L1RPCUrl                  string
	L1ChainID                 uint64
	L1ChainName               string
	L1UnstableBlocksDepth     uint64
	L1BlockScanBatchSizeLimit uint64

	L2RPCUrl                  string
	L2ChainID                 uint64
	L2ChainName               string
	L2UnstableBlocksDepth     uint64
	L2BlockScanBatchSizeLimit uint64

	SystemConfigAddr string

	DBConnStr     string
	DBMaxOpenConn int
	DBMaxIdleConn int
	DBMaxIdleTime string
}
