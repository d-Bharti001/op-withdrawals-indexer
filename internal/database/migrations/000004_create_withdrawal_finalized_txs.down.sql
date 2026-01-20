ALTER TABLE IF EXISTS withdrawal_finalized_txs
DROP CONSTRAINT fk_withdrawal;

ALTER TABLE IF EXISTS withdrawal_finalized_txs
DROP CONSTRAINT unique_withdrawal_finalized;

DROP TABLE IF EXISTS withdrawal_finalized_txs;
