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

	Nonce    *big.Int       `json:"withdrawal_nonce"`
	Sender   common.Address `json:"withdrawal_sender"`
	Target   common.Address `json:"withdrawal_target"`
	Value    *big.Int       `json:"withdrawal_value"`
	GasLimit *big.Int       `json:"withdrawal_gas_limit"`
	Data     hexutil.Bytes  `json:"withdrawal_data"`

	TxHash         common.Hash `json:"tx_hash"`
	BlockNumber    uint64      `json:"block_number"`
	BlockHash      common.Hash `json:"block_hash"`
	BlockTimestamp uint64      `json:"block_timestamp"`
}

// Database specific model
type WithdrawalDBRow struct {
	Hash    dbtypes.Hash
	ChainID uint64

	Nonce    dbtypes.U256
	Sender   dbtypes.Address
	Target   dbtypes.Address
	Value    dbtypes.U256
	GasLimit dbtypes.U256
	Data     dbtypes.Bytes

	TxHash         dbtypes.Hash
	BlockNumber    uint64
	BlockHash      dbtypes.Hash
	BlockTimestamp uint64
}

func (m *Withdrawal) ToDBRow() *WithdrawalDBRow {
	return &WithdrawalDBRow{
		Hash:           dbtypes.Hash(m.Hash),
		ChainID:        m.ChainID,
		Nonce:          dbtypes.NewU256(m.Nonce),
		Sender:         dbtypes.Address(m.Sender),
		Target:         dbtypes.Address(m.Target),
		Value:          dbtypes.NewU256(m.Value),
		GasLimit:       dbtypes.NewU256(m.GasLimit),
		Data:           dbtypes.Bytes(m.Data),
		TxHash:         dbtypes.Hash(m.TxHash),
		BlockNumber:    m.BlockNumber,
		BlockHash:      dbtypes.Hash(m.BlockHash),
		BlockTimestamp: m.BlockTimestamp,
	}
}

func (r *WithdrawalDBRow) ToDomainModel() *Withdrawal {
	return &Withdrawal{
		Hash:           r.Hash.Common(),
		ChainID:        r.ChainID,
		Nonce:          r.Nonce.BigInt(),
		Sender:         r.Sender.Common(),
		Target:         r.Target.Common(),
		Value:          r.Value.BigInt(),
		GasLimit:       r.GasLimit.BigInt(),
		Data:           r.Data.HexUtilBytes(),
		TxHash:         r.TxHash.Common(),
		BlockNumber:    r.BlockNumber,
		BlockHash:      r.BlockHash.Common(),
		BlockTimestamp: r.BlockTimestamp,
	}
}
