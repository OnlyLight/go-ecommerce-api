-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_info_9999` (
    user_id BIGINT(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
    user_account VARCHAR(255) NOT NULL COMMENT 'User account',
    user_nickname VARCHAR(255) COMMENT 'User nickname',
    user_avatar VARCHAR(255) COMMENT 'User avatar',
    user_state TINYINT UNSIGNED NOT NULL DEFAULT 2 COMMENT 'User state: 0: Locked, 1: Activated, 2: Not Activated',
    user_mobile VARCHAR(20) COMMENT 'User mobile',

    user_gender TINYINT UNSIGNED DEFAULT 0 COMMENT 'User gender: 0-Secret, 1-Male, 2-Female',
    user_birthday DATE COMMENT 'User birthday',
    user_email VARCHAR(255) COMMENT 'User email',

    user_is_authentication TINYINT UNSIGNED DEFAULT 0 NOT NULL COMMENT 'User is authentication: 0: Not authenticated, 1: Pending, 2: Authenticated',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time',

    -- Index for optimize query
    UNIQUE KEY `unique_user_account` (`user_account`),
    INDEX `idx_user_mobile` (`user_mobile`),
    INDEX `idx_user_email` (`user_email`),
    INDEX `idx_user_state` (`user_state`),
    INDEX `idx_user_is_authentication` (`user_is_authentication`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='pre_go_acc_user_info_9999';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_info_9999`;
-- +goose StatementEnd
