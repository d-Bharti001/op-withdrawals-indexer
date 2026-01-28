package models

import (
	"math/big"
	"op-withdrawals-indexer/internal/database/dbtypes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Domain / application specific model
type Withdrawal struct {
	Hash    common.Hash `json:"withdrawal_hash"`
	ChainID uint64      `json:"chain_id"`

	Sender common.Address `json:"withdrawal_sender"`
	Target common.Address `json:"withdrawal_target"`
	Data   hexutil.Bytes  `json:"withdrawal_data"`
	Value  *big.Int       `json:"withdrawal_value"`

	TxHash         common.Hash    `json:"tx_hash"`
	TxCaller       common.Address `json:"tx_caller"`
	BlockNumber    uint64         `json:"block_number"`
	BlockHash      common.Hash    `json:"block_hash"`
	BlockTimestamp uint64         `json:"block_timestamp"`
}

// Database specific model
type WithdrawalDBRow struct {
	Hash    dbtypes.Hash
	ChainID uint64

	Sender dbtypes.Address
	Target dbtypes.Address
	Data   dbtypes.Bytes
	Value  dbtypes.U256

	TxHash         dbtypes.Hash
	TxCaller       dbtypes.Address
	BlockNumber    uint64
	BlockHash      dbtypes.Hash
	BlockTimestamp uint64
}

func (m *Withdrawal) ToDBRow() *WithdrawalDBRow {
	return &WithdrawalDBRow{
		Hash:           dbtypes.Hash(m.Hash),
		ChainID:        m.ChainID,
		Sender:         dbtypes.Address(m.Sender),
		Target:         dbtypes.Address(m.Target),
		Data:           dbtypes.Bytes(m.Data),
		Value:          dbtypes.NewU256(m.Value),
		TxHash:         dbtypes.Hash(m.TxHash),
		TxCaller:       dbtypes.Address(m.TxCaller),
		BlockNumber:    m.BlockNumber,
		BlockHash:      dbtypes.Hash(m.BlockHash),
		BlockTimestamp: m.BlockTimestamp,
	}
}

func (r *WithdrawalDBRow) ToDomainModel() *Withdrawal {
	return &Withdrawal{
		Hash:           r.Hash.Common(),
		ChainID:        r.ChainID,
		Sender:         r.Sender.Common(),
		Target:         r.Target.Common(),
		Data:           r.Data.HexUtilBytes(),
		Value:          r.Value.BigInt(),
		TxHash:         r.TxHash.Common(),
		TxCaller:       r.TxCaller.Common(),
		BlockNumber:    r.BlockNumber,
		BlockHash:      r.BlockHash.Common(),
		BlockTimestamp: r.BlockTimestamp,
	}
}
