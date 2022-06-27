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


-- +migrate Down
DROP TABLE `users`;
DROP TABLE `departments`;
DROP TABLE `department_users`;