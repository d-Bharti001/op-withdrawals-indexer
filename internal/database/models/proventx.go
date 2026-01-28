package models

import (
	"op-withdrawals-indexer/internal/database/dbtypes"

	"github.com/ethereum/go-ethereum/common"
)

// Domain / application specific model
type WithdrawalProvenTx struct {
	WithdrawalHash    common.Hash `json:"withdrawal_hash"`
	WithdrawalChainID uint64      `json:"withdrawal_chain_id"`

	ProofSubmitter common.Address `json:"proof_submitter"`

	TxHash         common.Hash    `json:"tx_hash"`
	TxChainID      uint64         `json:"tx_chain_id"`
	TxCaller       common.Address `json:"tx_caller"`
	BlockNumber    uint64         `json:"block_number"`
	BlockHash      common.Hash    `json:"block_hash"`
	BlockTimestamp uint64         `json:"block_timestamp"`
}

// Database specific model
type WithdrawalProvenTxDBRow struct {
	WithdrawalHash    dbtypes.Hash
	WithdrawalChainID uint64

	ProofSubmitter dbtypes.Address

	TxHash         dbtypes.Hash
	TxChainID      uint64
	TxCaller       dbtypes.Address
	BlockNumber    uint64
	BlockHash      dbtypes.Hash
	BlockTimestamp uint64
}

func (m *WithdrawalProvenTx) ToDBRow() *WithdrawalProvenTxDBRow {
	return &WithdrawalProvenTxDBRow{
		WithdrawalHash:    dbtypes.Hash(m.WithdrawalHash),
		WithdrawalChainID: m.WithdrawalChainID,
		ProofSubmitter:    dbtypes.Address(m.ProofSubmitter),
		TxHash:            dbtypes.Hash(m.TxHash),
		TxChainID:         m.TxChainID,
		TxCaller:          dbtypes.Address(m.TxCaller),
		BlockNumber:       m.BlockNumber,
		BlockHash:         dbtypes.Hash(m.BlockHash),
		BlockTimestamp:    m.BlockTimestamp,
	}
}

func (r *WithdrawalProvenTxDBRow) ToDomainModel() *WithdrawalProvenTx {
	return &WithdrawalProvenTx{
		WithdrawalHash:    r.WithdrawalHash.Common(),
		WithdrawalChainID: r.WithdrawalChainID,
		ProofSubmitter:    r.ProofSubmitter.Common(),
		TxHash:            r.TxHash.Common(),
		TxChainID:         r.TxChainID,
		TxCaller:          r.TxCaller.Common(),
		BlockNumber:       r.BlockNumber,
		BlockHash:         r.BlockHash.Common(),
		BlockTimestamp:    r.BlockTimestamp,
	}
}
