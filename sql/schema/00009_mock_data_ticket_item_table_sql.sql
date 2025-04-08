-- +goose Up
-- +goose StatementBegin
INSERT INTO `ticket_item` (`name`, `description`, `stock_initial`, `stock_available`, `is_stock_prepared`, `price_original`, `price_flash`, `sale_start_time`, `sale_end_time`, `status`, `activity_id`, `updated_at`, `created_at`) 
VALUES 
('VIP Seat - Concert A', 'Front row VIP seating for Concert A', 100, 80, 1, 200000, 150000, '2025-02-25 10:00:00', '2025-03-01 10:00:00', 1, 1, NOW(), NOW()),
('General Entry - Tech Conference', 'Standard entry ticket for Tech Conference', 500, 450, 1, 50000, 40000, '2025-04-10 08:00:00', '2025-04-15 09:00:00', 1, 2, NOW(), NOW()),
('Premium Ticket - Sports Match', 'Exclusive seating with hospitality', 200, 180, 1, 150000, 120000, '2025-05-01 12:00:00', '2025-05-10 15:00:00', 0, 3, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `ticket_item`
-- +goose StatementEnd
