CREATE TABLE IF NOT EXISTS withdrawal_finalized_txs (
    withdrawal_hash CHAR(66) NOT NULL,
    withdrawal_chain_id BIGINT NOT NULL,

    -- event WithdrawalFinalized(withdrawalHash, success)
    success BOOLEAN NOT NULL,

    tx_hash CHAR(66) NOT NULL CHECK (tx_hash = LOWER(tx_hash)),
    tx_chain_id BIGINT NOT NULL,
    block_number BIGINT NOT NULL,
    block_hash CHAR(66) NOT NULL CHECK (block_hash = LOWER(block_hash)),
    block_timestamp BIGINT NOT NULL,

    CONSTRAINT fk_withdrawal
        FOREIGN KEY (withdrawal_chain_id, withdrawal_hash)
        REFERENCES withdrawals (chain_id, withdrawal_hash),

    CONSTRAINT unique_withdrawal_finalized
        UNIQUE (withdrawal_chain_id, withdrawal_hash, tx_chain_id)
);
