CREATE TABLE IF NOT EXISTS withdrawal_info (
    withdrawal_hash CHAR(66) NOT NULL,
    chain_id BIGINT NOT NULL,

    decoded_action VARCHAR(100) NOT NULL DEFAULT '',

    token_address CHAR(42),
    from_address CHAR(42) CHECK (from_address = LOWER(from_address)),
    to_address CHAR(42) CHECK (to_address = LOWER(to_address)),
    token_id NUMERIC(78, 0),
    amount NUMERIC(78, 0),

    CONSTRAINT fk_withdrawal
        FOREIGN KEY (chain_id, withdrawal_hash)
        REFERENCES withdrawals (chain_id, withdrawal_hash),
    
    CONSTRAINT fk_token
        FOREIGN KEY (chain_id, token_address)
        REFERENCES tokens (chain_id, token_address),

    CONSTRAINT unique_withdrawal_info
        UNIQUE (chain_id, withdrawal_hash)
);

ALTER TABLE IF EXISTS withdrawal_info
ADD CONSTRAINT check_decoded_action
CHECK (decoded_action IN (
    'eth_transfer',
    'erc20_transfer',
    'erc721_transfer',
    'erc1155_transfer'
));
