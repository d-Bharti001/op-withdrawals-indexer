package blockchain

import (
	"op-withdrawals-indexer/internal/contracts/predeploys"

	"github.com/ethereum/go-ethereum/common"
)

type L1ContractAddresses struct {
	optimismPortalAddr         common.Address
	l1CrossDomainMessengerAddr common.Address
	l1StandardBridgeAddr       common.Address
	l1ERC721BridgeAddr         common.Address
}

func (contracts *L1ContractAddresses) OptimismPortalAddr() common.Address {
	return contracts.optimismPortalAddr
}
func (contracts *L1ContractAddresses) L1CrossDomainMessengerAddr() common.Address {
	return contracts.l1CrossDomainMessengerAddr
}
func (contracts *L1ContractAddresses) L1StandardBridgeAddr() common.Address {
	return contracts.l1StandardBridgeAddr
}
func (contracts *L1ContractAddresses) L1ERC721BridgeAddr() common.Address {
	return contracts.l1ERC721BridgeAddr
}

type L2ContractAddresses struct{}

func (L2ContractAddresses) L2ToL1MessagePasserAddr() common.Address {
	return predeploys.L2ToL1MessagePasserAddr
}

func (L2ContractAddresses) L2CrossDomainMessengerAddr() common.Address {
	return predeploys.L2CrossDomainMessengerAddr
}

func (L2ContractAddresses) L2StandardBridgeAddr() common.Address {
	return predeploys.L2StandardBridgeAddr
}

func (L2ContractAddresses) L2ERC721BridgeAddr() common.Address {
	return predeploys.L2ERC721BridgeAddr
}
