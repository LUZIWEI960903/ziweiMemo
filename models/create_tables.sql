DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20) NOT NULL,
    `username`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       varchar(64) COLLATE utf8mb4_general_ci,
    `gender`      tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `task_id`     bigint(20) NOT NULL,
    `title`       varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `user_id`     bigint(20) NOT NULL COMMENT '作者id',
    `status`      tinyint(4) NOT NULL DEFAULT '0' COMMENT '备忘录状态',
    `content`     varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `start_time`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
    `end_time`    timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_task_id` (`task_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

ALTER TABLE `task` ADD COLUMN `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除';