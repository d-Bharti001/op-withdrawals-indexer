package predeploys

import "github.com/ethereum/go-ethereum/common"

var (
	L2ToL1MessagePasserAddr    common.Address = common.HexToAddress("0x4200000000000000000000000000000000000016")
	L2CrossDomainMessengerAddr common.Address = common.HexToAddress("0x4200000000000000000000000000000000000007")
	L2StandardBridgeAddr       common.Address = common.HexToAddress("0x4200000000000000000000000000000000000010")
	L2ERC721BridgeAddr         common.Address = common.HexToAddress("0x4200000000000000000000000000000000000014")
)
