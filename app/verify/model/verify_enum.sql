CREATE TABLE `verify_enum`
(
    `id`     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `verify_type` varchar(64) NOT NULL COMMENT '验证类型',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `is_delete` INT NOT NULL DEFAULT '0' COMMENT '删除标识 0 未删除 1 已删除',
    `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='verify enum'
