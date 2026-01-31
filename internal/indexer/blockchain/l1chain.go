package blockchain

import (
	"context"
	"math/big"
	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/contracts/signatures"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type L1Chain struct {
	Chain

	iOptimismPortal2 *bindings.IOptimismPortal2
}

func (c *L1Chain) PortalWithdrawalLogs(ctx context.Context, portalAddrs []common.Address, fromBlock, toBlock uint64) ([]OptimismPortalRelevantWithdrawalEvent, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: portalAddrs,
		Topics: [][]common.Hash{
			{
				signatures.IOptimismPortal2WithdrawalProvenExtension1EventSig,
				signatures.IOptimismPortal2WithdrawalFinalizedEventSig,
			},
		},
	}

	ctx, cancel := context.WithTimeout(ctx, DefaultFilterLogsTimeout)
	defer cancel()

	logs, err := c.rpcClient.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	parsedLogs := make([]OptimismPortalRelevantWithdrawalEvent, len(logs))

	for i, log := range logs {
		switch log.Topics[0] {
		case signatures.IOptimismPortal2WithdrawalProvenExtension1EventSig:
			parsed, err := c.iOptimismPortal2.UnpackWithdrawalProvenExtension1Event(&log)
			if err != nil {
				return nil, signatures.IOptimismPortal2WithdrawalProvenExtension1EventUnpackError
			}
			parsedLogs[i] = &WithdrawalProvenEvent{parsed}

		case signatures.IOptimismPortal2WithdrawalFinalizedEventSig:
			parsed, err := c.iOptimismPortal2.UnpackWithdrawalFinalizedEvent(&log)
			if err != nil {
				return nil, signatures.IOptimismPortal2WithdrawalFinalizedEventUnpackError
			}
			parsedLogs[i] = &WithdrawalFinalizedEvent{parsed}

		default:
			return nil, signatures.EventUnpackError
		}
	}

	return parsedLogs, nil
}
