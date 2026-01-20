CREATE TABLE IF NOT EXISTS withdrawal_proven_txs (
    withdrawal_hash CHAR(66) NOT NULL,
    withdrawal_chain_id BIGINT NOT NULL,

    -- event WithdrawalProvenExtension1(withdrawalHash, proofSubmitter)
    proof_submitter CHAR(42) NOT NULL CHECK (proof_submitter = LOWER(proof_submitter)),

    tx_hash CHAR(66) NOT NULL CHECK (tx_hash = LOWER(tx_hash)),
    tx_chain_id BIGINT NOT NULL,
    tx_caller CHAR(42) NOT NULL CHECK (tx_caller = LOWER(tx_caller)),
    block_number BIGINT NOT NULL,
    block_hash CHAR(66) NOT NULL CHECK (block_hash = LOWER(block_hash)),
    block_timestamp BIGINT NOT NULL,

    CONSTRAINT fk_withdrawal
        FOREIGN KEY (withdrawal_chain_id, withdrawal_hash)
        REFERENCES withdrawals (chain_id, withdrawal_hash),

    CONSTRAINT unique_withdrawal_proven
        UNIQUE (withdrawal_chain_id, withdrawal_hash, proof_submitter, tx_hash)
);
