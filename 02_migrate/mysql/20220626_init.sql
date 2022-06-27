
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `account`    varchar(191) DEFAULT NULL COMMENT '账号',
    `password`   varchar(191) DEFAULT NULL COMMENT '密码',
    `username`   varchar(191) DEFAULT NULL COMMENT '昵称',
    `phone`      varchar(16)  DEFAULT NULL COMMENT '手机号',
    `avatar`     varchar(191) DEFAULT NULL COMMENT '头像',
    `email`      varchar(191) DEFAULT NULL COMMENT '头像',
    `created_at` datetime     DEFAULT NULL,
    `updated_at` datetime     DEFAULT NULL,
    `deleted_at` datetime     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



-- +migrate Down
DROP TABLE `users`;