ALTER TABLE IF EXISTS withdrawal_info
DROP CONSTRAINT fk_withdrawal;

ALTER TABLE IF EXISTS withdrawal_info
DROP CONSTRAINT unique_withdrawal_info;

ALTER TABLE IF EXISTS withdrawal_info
DROP CONSTRAINT check_decoded_action;

DROP TABLE IF EXISTS withdrawal_info;
