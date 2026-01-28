package models

import (
	"op-withdrawals-indexer/internal/database/dbtypes"

	"github.com/ethereum/go-ethereum/common"
)

// Domain / application specific model
type WithdrawalFinalizedTx struct {
	WithdrawalHash    common.Hash `json:"withdrawal_hash"`
	WithdrawalChainID uint64      `json:"withdrawal_chain_id"`

	Success bool `json:"success"`

	TxHash         common.Hash    `json:"tx_hash"`
	TxChainID      uint64         `json:"tx_chain_id"`
	TxCaller       common.Address `json:"tx_caller"`
	BlockNumber    uint64         `json:"block_number"`
	BlockHash      common.Hash    `json:"block_hash"`
	BlockTimestamp uint64         `json:"block_timestamp"`
}

// Database specific model
type WithdrawalFinalizedTxDBRow struct {
	WithdrawalHash    dbtypes.Hash
	WithdrawalChainID uint64

	Success bool

	TxHash         dbtypes.Hash
	TxChainID      uint64
	TxCaller       dbtypes.Address
	BlockNumber    uint64
	BlockHash      dbtypes.Hash
	BlockTimestamp uint64
}

func (m *WithdrawalFinalizedTx) ToDBRow() *WithdrawalFinalizedTxDBRow {
	return &WithdrawalFinalizedTxDBRow{
		WithdrawalHash:    dbtypes.Hash(m.WithdrawalHash),
		WithdrawalChainID: m.WithdrawalChainID,
		Success:           m.Success,
		TxHash:            dbtypes.Hash(m.TxHash),
		TxChainID:         m.TxChainID,
		TxCaller:          dbtypes.Address(m.TxCaller),
		BlockNumber:       m.BlockNumber,
		BlockHash:         dbtypes.Hash(m.BlockHash),
		BlockTimestamp:    m.BlockTimestamp,
	}
}

func (r *WithdrawalFinalizedTxDBRow) ToDomainModel() *WithdrawalFinalizedTx {
	return &WithdrawalFinalizedTx{
		WithdrawalHash:    r.WithdrawalHash.Common(),
		WithdrawalChainID: r.WithdrawalChainID,
		Success:           r.Success,
		TxHash:            r.TxHash.Common(),
		TxChainID:         r.TxChainID,
		TxCaller:          r.TxCaller.Common(),
		BlockNumber:       r.BlockNumber,
		BlockHash:         r.BlockHash.Common(),
		BlockTimestamp:    r.BlockTimestamp,
	}
}
