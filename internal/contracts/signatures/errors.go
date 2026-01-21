package signatures

import "errors"

var IL2ToL1MessagePasserMessagePassedEventUnpackError = errors.New("failed to parse IL2ToL1MessagePasse MessagePassed event")

var IOptimismPortal2WithdrawalProvenExtension1EventUnpackError = errors.New("failed to parse IOptimismPortal2 WithdrawalProvenExtension1 event")

var IOptimismPortal2WithdrawalFinalizedEventUnpackError = errors.New("failed to parse IOptimismPortal2 WithdrawalFinalized event")

var EventUnpackError = errors.New("failed to parse or recognize event log")
