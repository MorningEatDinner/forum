create database community;
use  community;
CREATE TABLE `communities` (
    `community_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '社区ID',
    `community_name` varchar(100) NOT NULL DEFAULT '' COMMENT '社区名称',
    `introduction` varchar(500) DEFAULT NULL COMMENT '社区介绍',
    `create_time` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_time` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`) COMMENT '社区名称唯一索引'
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_unicode_ci 
  COMMENT='社区信息表';