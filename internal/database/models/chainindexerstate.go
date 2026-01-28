package models

// Domain / application specific model
type ChainIndexerState struct {
	ChainKey                string `json:"chain_key"`
	LatestSyncedBlockNumber uint64 `json:"latest_synced_block_number"`
}

// Database specific model
// (no conversion needed as it uses the native data types)
type ChainIndexerStateDBRow struct {
	ChainIndexerState
}
