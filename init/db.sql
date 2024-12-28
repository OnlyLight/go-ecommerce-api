ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root1234';
FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS `users` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(191) NOT NULL,
    `email` VARCHAR(191) NOT NULL,
    PRIMARY KEY (`id`)
)  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
