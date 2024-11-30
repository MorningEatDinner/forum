CREATE TABLE posts (
    post_id      BIGINT AUTO_INCREMENT COMMENT '帖子ID',
    author_id    BIGINT NOT NULL DEFAULT 0 COMMENT '作者ID',
    community_id BIGINT NOT NULL DEFAULT 0 COMMENT '社区ID',
    title        VARCHAR(255) NOT NULL DEFAULT '' COMMENT '帖子标题',
    content      LONGTEXT NOT NULL COMMENT '帖子内容',
    create_time  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_time DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    PRIMARY KEY(post_id),
    INDEX idx_community (community_id, create_time) COMMENT '社区帖子索引',
    INDEX idx_author (author_id, create_time) COMMENT '作者帖子索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='帖子表';

CREATE TABLE posts_stats (
     id            BIGINT AUTO_INCREMENT COMMENT '记录ID',
     post_id       BIGINT NOT NULL COMMENT '帖子ID',
     upvotes       INT NOT NULL DEFAULT 0 COMMENT '点赞数',
     downvotes     INT NOT NULL DEFAULT 0 COMMENT '点踩数',
     create_time   DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
     updated_time  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
     PRIMARY KEY(id),
     UNIQUE KEY uk_post_id (post_id) COMMENT '帖子ID唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='帖子投票统计表';
