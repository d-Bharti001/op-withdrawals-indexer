CREATE TABLE IF NOT EXISTS chains (
    chain_id BIGINT PRIMARY KEY,
    chain_name VARCHAR(255) NOT NULL,

    -- the lower layer chain that this chain settles to
    source_chain_id BIGINT REFERENCES chains(chain_id)
);
