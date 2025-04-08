CREATE DATABASE IF NOT EXISTS `shopdevgo`
DEFAULT CHARSET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- ticket table
CREATE TABLE IF NOT EXISTS `shopdevgo`.`ticket` (
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

-- ticket detail (item) table
CREATE TABLE IF NOT EXISTS `shopdevgo`.`ticket_item` (
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

-- Mock Data
-- Insert mock data into ticket table
INSERT INTO `shopdevgo`.`ticket` (`name`, `description`, `start_time`, `end_time`, `status`, `updated_at`, `created_at`) 
VALUES 
('Concert A', 'Music concert event', '2025-03-01 10:00:00', '2025-03-01 22:00:00', 1, NOW(), NOW()),
('Tech Conference', 'Annual technology conference', '2025-04-15 09:00:00', '2025-04-16 18:00:00', 1, NOW(), NOW()),
('Sports Match', 'Football league final', '2025-05-10 15:00:00', '2025-05-10 20:00:00', 0, NOW(), NOW());

-- Insert mock data into ticket_item table
INSERT INTO `shopdevgo`.`ticket_item` (`name`, `description`, `stock_initial`, `stock_available`, `is_stock_prepared`, `price_original`, `price_flash`, `sale_start_time`, `sale_end_time`, `status`, `activity_id`, `updated_at`, `created_at`) 
VALUES 
('VIP Seat - Concert A', 'Front row VIP seating for Concert A', 100, 80, 1, 200000, 150000, '2025-02-25 10:00:00', '2025-03-01 10:00:00', 1, 1, NOW(), NOW()),
('General Entry - Tech Conference', 'Standard entry ticket for Tech Conference', 500, 450, 1, 50000, 40000, '2025-04-10 08:00:00', '2025-04-15 09:00:00', 1, 2, NOW(), NOW()),
('Premium Ticket - Sports Match', 'Exclusive seating with hospitality', 200, 180, 1, 150000, 120000, '2025-05-01 12:00:00', '2025-05-10 15:00:00', 0, 3, NOW(), NOW());