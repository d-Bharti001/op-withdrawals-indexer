package blockchain

import (
	"context"
	"op-withdrawals-indexer/internal/contracts/bindings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ChainProviderCharacteristics interface {
	ID() uint64
	Name() string
	LatestBlockNumber(ctx context.Context) (uint64, error)
	BlockConfirmationDepth() uint64
	PollingInterval() time.Duration
}

type L2ChainProvider interface {
	ChainProviderCharacteristics

	L1ChainID() uint64
	L1StartBlockNumber() uint64

	// Should return MessagePassed events from the L2ToL1MessagePasser contract
	MessagePassedLogs(ctx context.Context, fromBlock, toBlock uint64) ([]bindings.IL2ToL1MessagePasserMessagePassed, error)
}

type L1ChainProvider interface {
	ChainProviderCharacteristics

	AttachedL2Chains() []L2ChainProvider

	// Should return WithdrawalProvenExtension1, WithdrawalFinalized events from OptimismPortal2 contracts
	PortalWithdrawalLogs(ctx context.Context, portalAddrs []common.Address, fromBlock, toBlock uint64) ([]types.Log, error)
}
