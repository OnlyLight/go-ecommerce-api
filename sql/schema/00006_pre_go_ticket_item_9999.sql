-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `ticket_item` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
  `name` VARCHAR(50) NOT NULL COMMENT 'ticket title',
  `description` TEXT COMMENT 'ticket description',
  `stock_initial` INT(11) NOT NULL DEFAULT 0 COMMENT 'Initial stock quantity (e.g 1000 tickets)',
  `stock_available` INT(11) NOT NULL DEFAULT 0 COMMENT 'Current available stock (e.g 900 tickets)',
  `is_stock_prepared` BOOLEAN NOT NULL DEFAULT 0 COMMENT 'Indicates if stock is pre-warmed (0/1)',
  `price_original` BIGINT(20) NOT NULL COMMENT 'Original ticket price',
  `price_flash` BIGINT(20) NOT NULL COMMENT 'Discounted price during flash sale',
  `sale_start_time` DATETIME NOT NULL COMMENT 'Flash sale start time',
  `sale_end_time` DATETIME NOT NULL COMMENT 'Flash sale end time',
  `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'Ticket status (e.g active/inactive)',
  `activity_id` BIGINT(20) NOT NULL COMMENT 'ID of associated activity', -- ID cua hoat dong lien quan den

  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last updated time',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  PRIMARY KEY (`id`),

  KEY `idx_end_time` (`sale_start_time`), -- Very high query runtime
  KEY `idx_start_time` (`sale_end_time`), -- Very high query runtime
  KEY `idx_status` (`status`) -- Very high query runtime
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='table for ticket details';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ticket_item`;
-- +goose StatementEnd
