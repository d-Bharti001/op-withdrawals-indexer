package blockchain

import (
	"time"
)

type ChainsInitConfig struct {
	L2RPC string
	L1RPC string

	L2ChainName string
	L1ChainName string

	L2ChainID uint64
	L1ChainID uint64

	SystemConfigAddr string

	L2BlockConfirmationDepth *uint64
	L1BlockConfirmationDepth *uint64
}

var DefaultL2ChainBlockConfDepth uint64 = 24
var DefaultL1ChainBlockConfDepth uint64 = 6

var DefaultL2PollingInterval = 5 * time.Second
var DefaultL1PollingInterval = 5 * time.Second

var DefaultRpcTimeout = 5 * time.Second
var DefaultFilterLogsTimeout = 5 * time.Second
