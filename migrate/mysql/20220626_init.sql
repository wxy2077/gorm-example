-- +migrate Up
# 用户表
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

# 部门表
CREATE TABLE IF NOT EXISTS `departments`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `dep_id`     int(11)    NOT NULL COMMENT '部门id',
    `title`      varchar(191) DEFAULT NULL COMMENT '名称',
    `parent_id`  varchar(191) DEFAULT NULL COMMENT '父id',
    `level`      varchar(191) DEFAULT NULL COMMENT '部门层级',
    `path`       varchar(16)  DEFAULT NULL COMMENT '部门id路径',
    `created_at` datetime     DEFAULT NULL,
    `updated_at` datetime     DEFAULT NULL,
    `deleted_at` datetime     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_departments_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


# 部门用户表 中间表
CREATE TABLE IF NOT EXISTS `department_users`
(
    `id`      bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` int(11)    NOT NULL COMMENT '用户id',
    `dep_id`  int(191)   NOT NULL COMMENT '部门id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO `users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'hanli', NULL, '韩立', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'mark', NULL, '马克', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'qi', NULL, '阿七', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'allen', NULL, '艾伦', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `users`(`id`, `account`, `password`, `username`, `phone`, `avatar`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'batman', NULL, '布鲁斯', NULL, NULL, NULL, NULL, NULL, NULL);

INSERT INTO `departments`(`id`, `dep_id`, `title`, `parent_id`, `level`, `path`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 1, '电影', '0', '0', '/1', NULL, NULL, NULL);
INSERT INTO `departments`(`id`, `dep_id`, `title`, `parent_id`, `level`, `path`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 2, '动漫', '0', '0', '/2', NULL, NULL, NULL);
INSERT INTO `departments`(`id`, `dep_id`, `title`, `parent_id`, `level`, `path`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 3, '国漫', '2', '1', '/2/3', NULL, NULL, NULL);
INSERT INTO `departments`(`id`, `dep_id`, `title`, `parent_id`, `level`, `path`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 4, '日漫', '2', '1', '/2/4', NULL, NULL, NULL);

INSERT INTO `department_users`(`id`, `user_id`, `dep_id`) VALUES (1, 1, 3);
INSERT INTO `department_users`(`id`, `user_id`, `dep_id`) VALUES (2, 2, 3);
INSERT INTO `department_users`(`id`, `user_id`, `dep_id`) VALUES (3, 3, 3);
INSERT INTO `department_users`(`id`, `user_id`, `dep_id`) VALUES (4, 4, 4);
INSERT INTO `department_users`(`id`, `user_id`, `dep_id`) VALUES (5, 5, 1);

-- +migrate Down
DROP TABLE `users`;
DROP TABLE `departments`;
DROP TABLE `department_users`;