package main

import (
	"errors"
	"op-withdrawals-indexer/internal/indexer/blockchain"
	indexer "op-withdrawals-indexer/internal/indexer/services"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v3"
)

var appConfig indexer.IndexerInitConfig

var appFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:        "l1RpcUrl",
		Usage:       "RPC URL for L1 chain",
		Required:    true,
		Sources:     cli.EnvVars("L1_RPC_URL"),
		Destination: &appConfig.L1RPCUrl,
	},
	&cli.Uint64Flag{
		Name:        "l1ChainId",
		Usage:       "Chain ID of L1 chain - must match the one returned by the rpc",
		Required:    true,
		Sources:     cli.EnvVars("L1_CHAIN_ID"),
		Destination: &appConfig.L1ChainID,
	},
	&cli.StringFlag{
		Name:        "l1ChainName",
		Usage:       "Name of L1 chain",
		Required:    true,
		Sources:     cli.EnvVars("L1_CHAIN_NAME"),
		Destination: &appConfig.L1ChainName,
	},
	&cli.Uint64Flag{
		Name:        "l1BlockConfirmationDepth",
		Usage:       "Depth from the latest block on L1 chain before which the blocks can be considered final",
		Sources:     cli.EnvVars("L1_BLOCK_CONFIRMATION_DEPTH"),
		Value:       blockchain.DefaultL1ChainBlockConfDepth,
		DefaultText: strconv.Itoa(int(blockchain.DefaultL1ChainBlockConfDepth)),
		Destination: &appConfig.L1BlockConfirmationDepth,
	},
	&cli.StringFlag{
		Name:        "l2RpcUrl",
		Usage:       "RPC URL for L2 chain",
		Required:    true,
		Sources:     cli.EnvVars("L2_RPC_URL"),
		Destination: &appConfig.L2RPCUrl,
	},
	&cli.Uint64Flag{
		Name:        "l2ChainId",
		Usage:       "Chain ID of L2 chain - must match the one returned by the rpc",
		Required:    true,
		Sources:     cli.EnvVars("L2_CHAIN_ID"),
		Destination: &appConfig.L2ChainID,
	},
	&cli.StringFlag{
		Name:        "l2ChainName",
		Usage:       "Name of L2 chain",
		Required:    true,
		Sources:     cli.EnvVars("L2_CHAIN_NAME"),
		Destination: &appConfig.L2ChainName,
	},
	&cli.Uint64Flag{
		Name:        "l2BlockConfirmationDepth",
		Usage:       "Depth from the latest block on L2 chain before which the blocks can be considered final",
		Sources:     cli.EnvVars("L2_BLOCK_CONFIRMATION_DEPTH"),
		Value:       blockchain.DefaultL2ChainBlockConfDepth,
		DefaultText: strconv.Itoa(int(blockchain.DefaultL2ChainBlockConfDepth)),
		Destination: &appConfig.L2BlockConfirmationDepth,
	},
	&cli.StringFlag{
		Name:        "systemConfigAddr",
		Usage:       "SystemConfig contract address (of the L2 chain) deployed on the L1 chain",
		Required:    true,
		Sources:     cli.EnvVars("SYSTEM_CONFIG_ADDR"),
		Validator:   AddressValidator(),
		Destination: &appConfig.SystemConfigAddr,
	},
	&cli.StringFlag{
		Name:        "dbConnString",
		Usage:       "Database connection string",
		Required:    true,
		Sources:     cli.EnvVars("DB_CONN_STR"),
		Destination: &appConfig.DBConnStr,
	},
	&cli.IntFlag{
		Name:        "dbMaxOpenConn",
		Usage:       "Database max open connections",
		Sources:     cli.EnvVars("DB_MAX_OPEN_CONN"),
		Value:       30,
		DefaultText: "30",
		Destination: &appConfig.DBMaxOpenConn,
	},
	&cli.IntFlag{
		Name:        "dbMaxIdleConn",
		Usage:       "Database max idle connections",
		Sources:     cli.EnvVars("DB_MAX_IDLE_CONN"),
		Value:       30,
		DefaultText: "30",
		Destination: &appConfig.DBMaxIdleConn,
	},
	&cli.StringFlag{
		Name:        "dbMaxIdleTime",
		Usage:       "Database max idle time",
		Sources:     cli.EnvVars("DB_MAX_IDLE_TIME"),
		Value:       "15m",
		DefaultText: "15m",
		Destination: &appConfig.DBMaxIdleTime,
	},
}

func AddressValidator() func(string) error {
	return func(s string) error {
		if common.IsHexAddress(s) {
			return nil
		}
		return errors.New("not an ethereum address")
	}
}
