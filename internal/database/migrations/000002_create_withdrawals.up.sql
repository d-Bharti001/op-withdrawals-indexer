CREATE TABLE IF NOT EXISTS withdrawals (
    withdrawal_hash CHAR(66) NOT NULL CHECK (withdrawal_hash = LOWER(withdrawal_hash)),
    chain_id BIGINT NOT NULL REFERENCES chains(chain_id),

    -- Values from MessagePassed event
    withdrawal_nonce NUMERIC(78, 0) NOT NULL,
    withdrawal_sender CHAR(42) NOT NULL CHECK (withdrawal_sender = LOWER(withdrawal_sender)),
    withdrawal_target CHAR(42) NOT NULL CHECK (withdrawal_target = LOWER(withdrawal_target)),
    withdrawal_value NUMERIC(78, 0) NOT NULL,
    withdrawal_gas_limit NUMERIC(78, 0) NOT NULL,
    withdrawal_data BYTEA NOT NULL,

    tx_hash CHAR(66) NOT NULL CHECK (tx_hash = LOWER(tx_hash)),
    block_number BIGINT NOT NULL,
    block_hash CHAR(66) NOT NULL CHECK (block_hash = LOWER(block_hash)),
    block_timestamp BIGINT NOT NULL,

    CONSTRAINT withdrawals_pk PRIMARY KEY (chain_id, withdrawal_hash)
);

CREATE INDEX IF NOT EXISTS
idx_withdrawal_initialization_block_timestamp_desc
ON withdrawals (block_timestamp DESC);
