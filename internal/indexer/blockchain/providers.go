package blockchain

import (
	"context"
	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

type ChainProvider interface {
	ChainID() uint64
	ChainName() string
	LatestBlockNumber(ctx context.Context) (uint64, error)
	TokenInfo(ctx context.Context, tokenAddr common.Address, typeHint models.TokenClass) (*models.Token, error)
	CloseConnection()
}

type L2ContractsProvider interface {
	L2ToL1MessagePasserAddr() common.Address
	L2CrossDomainMessengerAddr() common.Address
	L2StandardBridgeAddr() common.Address
	L2ERC721BridgeAddr() common.Address
}

type L1ContractsProvider interface {
	OptimismPortalAddr() common.Address
	L1CrossDomainMessengerAddr() common.Address
	L1StandardBridgeAddr() common.Address
	L1ERC721BridgeAddr() common.Address
}

type L2ChainProvider interface {
	ChainProvider
	L2ContractsProvider
	L1ContractsProvider

	L1ChainID() uint64
	L1StartBlockNumber() uint64

	// Should return MessagePassed events from the L2ToL1MessagePasser contract
	MessagePassedLogs(ctx context.Context, fromBlock, toBlock uint64) ([]*bindings.IL2ToL1MessagePasserMessagePassed, error)
}

type L1ChainProvider interface {
	ChainProvider

	// Should return WithdrawalProvenExtension1, WithdrawalFinalized events from OptimismPortal2 contracts
	PortalWithdrawalLogs(ctx context.Context, portalAddrs []common.Address, fromBlock, toBlock uint64) ([]OptimismPortalRelevantWithdrawalEvent, error)
}
