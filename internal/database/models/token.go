package models

import (
	"op-withdrawals-indexer/internal/database/dbtypes"

	"github.com/ethereum/go-ethereum/common"
)

// Domain / application specific model
type Token struct {
	Address  common.Address `json:"token_address"`
	ChainID  uint64         `json:"chain_id"`
	Class    string         `json:"token_class"`
	Name     string         `json:"token_name"`
	Symbol   string         `json:"token_symbol"`
	Decimals uint64         `json:"token_decimals"`
}

// Database specific model
type TokenDBRow struct {
	Address  dbtypes.Address
	ChainID  uint64
	Class    string
	Name     string
	Symbol   string
	Decimals uint64
}

type TokenClass string

const (
	ERC20Token   TokenClass = "ERC20"
	ERC721Token  TokenClass = "ERC721"
	ERC1155Token TokenClass = "ERC1155"
)

func (m *Token) ToDBRow() *TokenDBRow {
	return &TokenDBRow{
		Address:  dbtypes.Address(m.Address),
		ChainID:  m.ChainID,
		Class:    m.Class,
		Name:     m.Name,
		Symbol:   m.Symbol,
		Decimals: m.Decimals,
	}
}

func (r *TokenDBRow) ToDomainModel() *Token {
	return &Token{
		Address:  r.Address.Common(),
		ChainID:  r.ChainID,
		Class:    r.Class,
		Name:     r.Name,
		Symbol:   r.Symbol,
		Decimals: r.Decimals,
	}
}
