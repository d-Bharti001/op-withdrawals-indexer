CREATE TABLE IF NOT EXISTS tokens (
    token_address CHAR(42) NOT NULL CHECK (token_address = LOWER(token_address)),
    chain_id BIGINT NOT NULL REFERENCES chains(chain_id),
    token_class VARCHAR(20) NOT NULL,
    token_name VARCHAR(255) NOT NULL,
    token_symbol VARCHAR(255) NOT NULL,
    token_decimals SMALLINT NOT NULL DEFAULT 0,

    CONSTRAINT id PRIMARY KEY (chain_id, token_address),

    CONSTRAINT check_token_class
        CHECK (token_class IN ('ERC20', 'ERC721', 'ERC1155'))
);
