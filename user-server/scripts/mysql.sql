CREATE DATABASE ResourceShare;

USE ResourceShare;

CREATE TABLE `users` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `name` VARCHAR(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
    `avatar` VARCHAR(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
    `sex` INT(11) DEFAULT NULL COMMENT '性别 1为男性，2为女性',
    `ip_info` TEXT DEFAULT NULL COMMENT 'ip信息',
    `status` INT(11) DEFAULT '0' COMMENT '使用状态 0.正常 1拉黑',
    `create_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `update_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
    `email` VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户邮箱',
    `password` VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户密码',
    `bio` TEXT DEFAULT NULL COMMENT '个人简介',
    `level` INT(11) DEFAULT 0 COMMENT '用户等级',
    `experience` INT(11) DEFAULT 0 COMMENT '用户经验',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_email` (`email`) USING BTREE,
    KEY `idx_create_time` (`create_time`) USING BTREE
) ENGINE=INNODB AUTO_INCREMENT=20000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';
