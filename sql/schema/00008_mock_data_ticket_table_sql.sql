-- +goose Up
-- +goose StatementBegin
INSERT INTO `ticket` (`name`, `description`, `start_time`, `end_time`, `status`, `updated_at`, `created_at`) 
VALUES 
('Concert A', 'Music concert event', '2025-03-01 10:00:00', '2025-03-01 22:00:00', 1, NOW(), NOW()),
('Tech Conference', 'Annual technology conference', '2025-04-15 09:00:00', '2025-04-16 18:00:00', 1, NOW(), NOW()),
('Sports Match', 'Football league final', '2025-05-10 15:00:00', '2025-05-10 20:00:00', 0, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `ticket`
-- +goose StatementEnd
