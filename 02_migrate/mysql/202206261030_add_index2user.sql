

-- +migrate Up
ALTER TABLE `users` ADD INDEX `idx_account` (`account`);

-- +migrate Down
ALTER TABLE `users` DROP INDEX `idx_account`;
