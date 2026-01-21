package blockchain

import (
	"time"
)

var DefaultL2ChainBlockConfDepth uint64 = 24
var DefaultL1ChainBlockConfDepth uint64 = 6

var DefaultRpcTimeout = 5 * time.Second
var DefaultFilterLogsTimeout = 10 * time.Second

type ChainInitConfig struct {
	ID                     uint64
	Name                   string
	RPCUrl                 string
	BlockConfirmationDepth *uint64
}

type L1ChainInitConfig struct {
	ChainInitConfig
}

type L2ChainInitConfig struct {
	ChainInitConfig

	L1RPCUrl  string
	L1ChainID uint64

	SystemConfigAddr string
}
