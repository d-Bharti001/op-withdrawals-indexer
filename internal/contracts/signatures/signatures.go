package signatures

import "github.com/ethereum/go-ethereum/crypto"

var IL2ToL1MessagePasserMessagePassedEventSig = crypto.Keccak256Hash(
	[]byte("MessagePassed(uint256,address,address,uint256,uint256,bytes,bytes32)"),
)

var IOptimismPortal2WithdrawalProvenExtension1EventSig = crypto.Keccak256Hash(
	[]byte("WithdrawalProvenExtension1(bytes32,address)"),
)

var IOptimismPortal2WithdrawalFinalizedEventSig = crypto.Keccak256Hash(
	[]byte("WithdrawalFinalized(bytes32,bool)"),
)
