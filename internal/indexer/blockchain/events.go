package blockchain

import "op-withdrawals-indexer/internal/contracts/bindings"

type OptimismPortalRelevantWithdrawalEvent interface {
	isEvent()
}

type WithdrawalProvenEvent struct {
	*bindings.IOptimismPortal2WithdrawalProvenExtension1
}

func (*WithdrawalProvenEvent) isEvent() {}

type WithdrawalFinalizedEvent struct {
	*bindings.IOptimismPortal2WithdrawalFinalized
}

func (*WithdrawalFinalizedEvent) isEvent() {}
