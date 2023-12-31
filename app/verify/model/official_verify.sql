CREATE TABLE `official_verify`
(
    `id`     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `verify_info` varchar(255) NOT NULL COMMENT '验证信息',
    `verify_type` varchar(64) NOT NULL COMMENT '验证类型',
    `social_name` varchar(64)  COMMENT '名称',
    `job_tiele` varchar(64)  COMMENT '职位',
    `creator` varchar(64)  COMMENT '创建人名称',
    `creator_id` INT NOT NULL DEFAULT '0'   COMMENT '创建人id',
    `is_pay` varchar(64)  COMMENT '是否接受付款',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `is_delete` INT NOT NULL DEFAULT '0' COMMENT '删除标识 0 未删除 1 已删除',
    `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_verify_info` (`verify_info`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='验证信息'
