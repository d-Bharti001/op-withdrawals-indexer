package main

import (
	"errors"
	"strconv"

	indexer "op-withdrawals-indexer/internal/indexer/services"

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
	&cli.StringFlag{
		Name:        "l1PollingInterval",
		Usage:       "L1 chain polling interval",
		Sources:     cli.EnvVars("L1_POLLING_INTERVAL"),
		Value:       indexer.DefaultL1PollingInterval,
		DefaultText: indexer.DefaultL1PollingInterval,
		Destination: &appConfig.L1PollingInterval,
	},
	&cli.Uint64Flag{
		Name:        "l1UnstableBlocksDepth",
		Usage:       "Number of blocks on the L1 chain, counted backwards from the latest block, that can be assumed to change state after being scanned. For depth = 0, only the nth scanned block is considered unstable. For depth = 1, nth and (n-1)th blocks are in the unstable range.",
		Sources:     cli.EnvVars("L1_UNSTABLE_BLOCKS_DEPTH"),
		Value:       indexer.DefaultL1ChainUnstableBlocksDepth,
		DefaultText: strconv.Itoa(int(indexer.DefaultL1ChainUnstableBlocksDepth)),
		Destination: &appConfig.L1UnstableBlocksDepth,
	},
	&cli.Uint64Flag{
		Name:        "l1BlockScanBatchSizeLimit",
		Usage:       "(Max) number of blocks to scan in a single batch on the L1 chain",
		Sources:     cli.EnvVars("L1_BLOCK_SCAN_BATCH_SIZE_LIMIT"),
		Value:       indexer.DefaultL1BlockScanBatchSizeLimit,
		DefaultText: strconv.Itoa(int(indexer.DefaultL1BlockScanBatchSizeLimit)),
		Destination: &appConfig.L1BlockScanBatchSizeLimit,
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
	&cli.StringFlag{
		Name:        "l2PollingInterval",
		Usage:       "L2 chain polling interval",
		Sources:     cli.EnvVars("L2_POLLING_INTERVAL"),
		Value:       indexer.DefaultL2PollingInterval,
		DefaultText: indexer.DefaultL2PollingInterval,
		Destination: &appConfig.L2PollingInterval,
	},
	&cli.Uint64Flag{
		Name:        "l2UnstableBlocksDepth",
		Usage:       "Number of blocks on the L2 chain, counted backwards from the latest block, that can be assumed to change state after being scanned. For depth = 0, only the nth scanned block is considered unstable. For depth = 1, nth and (n-1)th blocks are in the unstable range.",
		Sources:     cli.EnvVars("L2_UNSTABLE_BLOCKS_DEPTH"),
		Value:       indexer.DefaultL2ChainUnstableBlocksDepth,
		DefaultText: strconv.Itoa(int(indexer.DefaultL2ChainUnstableBlocksDepth)),
		Destination: &appConfig.L2UnstableBlocksDepth,
	},
	&cli.Uint64Flag{
		Name:        "l2BlockScanBatchSizeLimit",
		Usage:       "(Max) number of blocks to scan in a single batch on the L2 chain",
		Sources:     cli.EnvVars("L2_BLOCK_SCAN_BATCH_SIZE_LIMIT"),
		Value:       indexer.DefaultL2BlockScanBatchSizeLimit,
		DefaultText: strconv.Itoa(int(indexer.DefaultL2BlockScanBatchSizeLimit)),
		Destination: &appConfig.L2BlockScanBatchSizeLimit,
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
