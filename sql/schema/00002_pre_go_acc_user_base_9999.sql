-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_base_9999` (
  user_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
  user_account VARCHAR(255) NOT NULL COMMENT 'User account',
  user_password VARCHAR(255) NOT NULL COMMENT 'User password',
  user_salt VARCHAR(255) NOT NULL COMMENT 'User salt',

  user_login_time TIMESTAMP NULL DEFAULT NULL COMMENT 'User login time',
  user_logout_time TIMESTAMP NULL DEFAULT NULL COMMENT 'User logout time',
  user_login_ip VARCHAR(45) NULL COMMENT 'User login IP',

  user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
  user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time',

  -- Index for optimize query
  UNIQUE KEY `unique_user_account` (`user_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='pre_go_acc_user_base_9999';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd
