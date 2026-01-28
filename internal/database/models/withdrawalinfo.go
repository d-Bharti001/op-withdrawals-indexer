package models

import (
	"math/big"
	"op-withdrawals-indexer/internal/database/dbtypes"

	"github.com/ethereum/go-ethereum/common"
)

// Domain / application specific model
type WithdrawalInformation struct {
	WithdrawalHash common.Hash `json:"withdrawal_hash"`
	ChainID        uint64      `json:"chain_id"`

	DecodedAction string `json:"decoded_action"`

	TokenAddress *common.Address `json:"token_address"`
	FromAddress  *common.Address `json:"from_address"`
	ToAddress    *common.Address `json:"to_address"`
	TokenID      *big.Int        `json:"token_id"`
	Amount       *big.Int        `json:"amount"`
}

// Database specific model
type WithdrawalInformationDBRow struct {
	WithdrawalHash dbtypes.Hash
	ChainID        uint64

	DecodedAction string

	TokenAddress dbtypes.NullableAddress
	FromAddress  dbtypes.NullableAddress
	ToAddress    dbtypes.NullableAddress
	TokenID      dbtypes.NullableU256
	Amount       dbtypes.NullableU256
}

type WithdrawalAction string

const (
	EthTransferAction     WithdrawalAction = "eth_transfer"
	ERC20TransferAction   WithdrawalAction = "erc20_transfer"
	ERC721TransferAction  WithdrawalAction = "erc721_transfer"
	ERC1155TransferAction WithdrawalAction = "erc1155_transfer"
)

func (m *WithdrawalInformation) ToDBRow() *WithdrawalInformationDBRow {
	return &WithdrawalInformationDBRow{
		WithdrawalHash: dbtypes.Hash(m.WithdrawalHash),
		ChainID:        m.ChainID,
		DecodedAction:  m.DecodedAction,
		TokenAddress:   dbtypes.NewNullableAddress(m.TokenAddress),
		FromAddress:    dbtypes.NewNullableAddress(m.FromAddress),
		ToAddress:      dbtypes.NewNullableAddress(m.ToAddress),
		TokenID:        dbtypes.NewNullableU256(m.TokenID),
		Amount:         dbtypes.NewNullableU256(m.Amount),
	}
}

func (r *WithdrawalInformationDBRow) ToDomainModel() *WithdrawalInformation {
	return &WithdrawalInformation{
		WithdrawalHash: r.WithdrawalHash.Common(),
		ChainID:        r.ChainID,
		DecodedAction:  r.DecodedAction,
		TokenAddress:   r.TokenAddress.Common(),
		FromAddress:    r.FromAddress.Common(),
		ToAddress:      r.ToAddress.Common(),
		TokenID:        r.TokenID.BigInt(),
		Amount:         r.Amount.BigInt(),
	}
}
