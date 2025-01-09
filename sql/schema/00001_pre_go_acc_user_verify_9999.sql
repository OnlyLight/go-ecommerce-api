-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_verify_9999` (
  verify_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
  verify_otp VARCHAR(255) NOT NULL COMMENT 'Verify otp',
  verify_type TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'User state: 1: EMAIL, 2: MOBILE',
  verify_key VARCHAR(255) NOT NULL COMMENT 'Verify key',
  verify_key_hash VARCHAR(255) NOT NULL COMMENT 'Verify key hash',
  is_verified BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Is verified: FALSE: Not Verify, TRUE: Verified',
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Is deleted: FALSE: Activing, TRUE: Deleted',

  verify_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
  verify_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time',

  -- Index for optimize query
  UNIQUE KEY `unique_verify_otp` (`verify_otp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='pre_go_acc_user_verify_9999';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
-- +goose StatementEnd
