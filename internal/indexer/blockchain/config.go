package blockchain

import (
	"time"
)

const DefaultRpcTimeout = 5 * time.Second
const DefaultFilterLogsTimeout = 10 * time.Second

type ChainInitConfig struct {
	ID     uint64
	Name   string
	RPCUrl string
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
