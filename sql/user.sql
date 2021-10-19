#
# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`
(
    `id`            bigint                           NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `gender`        tinyint(1)                       NOT NULL DEFAULT '0' COMMENT '1男2女0未知',
    `name`          varchar(32) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '用户名',
    `nick_name`     varchar(32) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '花名,绰号等',
    `password`      varchar(32) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '密码',
    `salt`          varchar(4) COLLATE utf8mb4_bin   NOT NULL DEFAULT '' COMMENT '混淆码',
    `mobile`        varchar(11) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '手机号码',
    `country_code`  varchar(8) COLLATE utf8mb4_bin   NOT NULL DEFAULT '' COMMENT '国家编码',
    `email`         varchar(64) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT 'email',
    `account`       varchar(64) COLLATE utf8mb4_bin  NOT NULL DEFAULT '' COMMENT '用户名',
    `register_type` tinyint(1)                                DEFAULT '0' COMMENT '注册类型：0 未知 1 手机注册 2 邮箱注册 3 账号注册',
    `avatar`        varchar(255) COLLATE utf8mb4_bin          DEFAULT '' COMMENT '自定义用户头像',
    `status`        tinyint(1)                                DEFAULT '0' COMMENT '用户状态：0 游客状态 1 普通状态 2 认证通过  3 冻结状态  4 注销状态',
    `sign_info`     varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '个性签名',
    `create_time`   timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_phone` (`mobile`)
) ENGINE = InnoDB AUTO_INCREMENT = 10000 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;