package blockchain

import (
	"context"
	"math/big"
	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/contracts/signatures"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type L2Chain struct {
	Chain
	L2ContractAddresses
	LinkedL1Details

	iL2ToL1MessagePasser *bindings.IL2ToL1MessagePasser
}

type LinkedL1Details struct {
	l1ChainId          uint64
	l1StartBlockNumber uint64
	L1ContractAddresses
}

func (l *LinkedL1Details) L1ChainID() uint64 {
	return l.l1ChainId
}

func (l *LinkedL1Details) L1StartBlockNumber() uint64 {
	return l.l1StartBlockNumber
}

func (c *L2Chain) MessagePassedLogs(ctx context.Context, fromBlock, toBlock uint64) ([]*bindings.IL2ToL1MessagePasserMessagePassed, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{c.L2ToL1MessagePasserAddr()},
		Topics: [][]common.Hash{
			{signatures.IL2ToL1MessagePasserMessagePassedEventSig},
		},
	}

	ctx, cancel := context.WithTimeout(ctx, DefaultFilterLogsTimeout)
	defer cancel()

	logs, err := c.rpcClient.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	parsedLogs := make([]*bindings.IL2ToL1MessagePasserMessagePassed, len(logs))

	for i, log := range logs {
		parsed, err := c.iL2ToL1MessagePasser.UnpackMessagePassedEvent(&log)
		if err != nil {
			return nil, signatures.IL2ToL1MessagePasserMessagePassedEventUnpackError
		}
		parsedLogs[i] = parsed
	}

	return parsedLogs, nil
}
