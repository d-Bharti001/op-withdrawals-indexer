DROP INDEX IF EXISTS idx_withdrawal_initialization_block_timestamp_desc;
ALTER TABLE IF EXISTS withdrawals DROP CONSTRAINT withdrawals_pk;
DROP TABLE IF EXISTS withdrawals;
