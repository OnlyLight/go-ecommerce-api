-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `ticket` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
  `name` VARCHAR(50) NOT NULL COMMENT 'ticket name',
  `description` TEXT COMMENT 'ticket description',
  `start_time` DATETIME NOT NULL COMMENT 'ticket sale start time',
  `end_time` DATETIME NOT NULL COMMENT 'ticket sale end time',
  `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'ticket sale activity status', -- 0: deactive, 1: active
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last updated time',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  PRIMARY KEY (`id`),

  KEY `idx_end_time` (`end_time`), -- Very high query runtime
  KEY `idx_start_time` (`start_time`), -- Very high query runtime
  KEY `idx_status` (`status`) -- Very high query runtime
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ticket table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ticket`;
-- +goose StatementEnd
