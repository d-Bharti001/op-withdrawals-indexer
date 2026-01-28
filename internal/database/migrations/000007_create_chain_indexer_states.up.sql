CREATE TABLE IF NOT EXISTS chain_indexer_states (
    chain_key VARCHAR(255) PRIMARY KEY,
    latest_synced_block BIGINT NOT NULL
);
