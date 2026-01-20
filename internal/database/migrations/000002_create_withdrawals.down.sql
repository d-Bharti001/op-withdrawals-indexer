DROP INDEX IF EXISTS idx_withdrawal_initialization_block_timestamp_desc;
ALTER TABLE IF EXISTS withdrawals DROP CONSTRAINT id;
DROP TABLE IF EXISTS withdrawals;
