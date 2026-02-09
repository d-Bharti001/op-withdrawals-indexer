package models

// Domain / application specific model
type Chain struct {
	ID            uint64  `json:"chain_id"`
	Name          string  `json:"chain_name"`
	SourceChainID *uint64 `json:"source_chain_id"`
}

// Database specific model
type ChainDBRow struct {
	ID            uint64
	Name          string
	SourceChainID *uint64
}

func (m *Chain) ToDBRow() *ChainDBRow {
	return &ChainDBRow{
		ID:            m.ID,
		Name:          m.Name,
		SourceChainID: m.SourceChainID,
	}
}

func (r *ChainDBRow) ToDomainModel() *Chain {
	return &Chain{
		ID:            r.ID,
		Name:          r.Name,
		SourceChainID: r.SourceChainID,
	}
}
