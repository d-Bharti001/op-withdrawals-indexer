package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

// TODO: delete this
// var (
// 	L2ToL1MessagePasserAddr    common.Address = common.HexToAddress("0x4200000000000000000000000000000000000016")
// 	L2CrossDomainMessengerAddr common.Address = common.HexToAddress("0x4200000000000000000000000000000000000007")
// 	L2StandardBridgeAddr       common.Address = common.HexToAddress("0x4200000000000000000000000000000000000010")
// 	L2ERC721BridgeAddr         common.Address = common.HexToAddress("0x4200000000000000000000000000000000000014")
// )

type L2Contracts struct {
	L2ToL1MessagePasser    *bind.BoundContract
	L2CrossDomainMessenger *bind.BoundContract
	L2StandardBridge       *bind.BoundContract
	L2ERC721Bridge         *bind.BoundContract
}

type L1Contracts struct {
	OptimismPortal         *bind.BoundContract
	L1CrossDomainMessenger *bind.BoundContract
	L1StandardBridge       *bind.BoundContract
	L1ERC721Bridge         *bind.BoundContract
}

// TODO: delete this
// type L1ContractAddresses struct {
// 	SystemConfigAddr           common.Address
// 	OptimismPortalAddr         common.Address
// 	L1CrossDomainMessengerAddr common.Address
// 	L1StandardBridgeAddr       common.Address
// 	L1ERC721BridgeAddr         common.Address
// }
