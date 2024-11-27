CREATE TABLE vote_record (
  vote_id      BIGINT AUTO_INCREMENT COMMENT '投票记录ID',
  post_id      BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子id', 
  user_id      BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  vote_type    TINYINT(4) NOT NULL DEFAULT 0 COMMENT '投票类型 0:赞成票 1:反对票',
  create_time  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  updated_time DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  
  PRIMARY KEY(vote_id),
  UNIQUE KEY uniq_post_user (post_id, user_id),
  INDEX idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='投票记录表';

CREATE TABLE vote_count (
  vote_count_id BIGINT AUTO_INCREMENT COMMENT '投票计数ID',
  post_id        BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子',
  agree_count   INT(11) NOT NULL DEFAULT 0 COMMENT '赞成票数',
  oppose_count  INT(11) NOT NULL DEFAULT 0 COMMENT '反对票数',
  create_time   DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  updated_time  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  
  PRIMARY KEY(vote_count_id),
  UNIQUE KEY uniq_post_id (post_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='投票计数表';