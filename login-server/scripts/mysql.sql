create Database ResourceShare;

USE ResourceShare;

CREATE TABLE
    `users` (
        `id` bigint (20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
        `name` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
        `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
        `sex` int (11) DEFAULT NULL COMMENT '性别 1为男性，2为女性',
        `ip_info` TEXT DEFAULT NULL COMMENT 'ip信息',
        `status` int (11) DEFAULT '0' COMMENT '使用状态 0.正常 1拉黑',
        `create_time` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
        `update_time` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
        `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户邮箱',
        `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户密码',
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE KEY `uniq_email` (`email`) USING BTREE,
        KEY `idx_create_time` (`create_time`) USING BTREE,
        KEY `idx_update_time` (`update_time`) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 20000 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC COMMENT = '用户表';
