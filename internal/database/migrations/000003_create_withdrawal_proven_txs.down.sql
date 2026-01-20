ALTER TABLE IF EXISTS withdrawal_proven_txs
DROP CONSTRAINT fk_withdrawal;

ALTER TABLE IF EXISTS withdrawal_proven_txs
DROP CONSTRAINT unique_withdrawal_proven;

DROP TABLE IF EXISTS withdrawal_proven_txs;
