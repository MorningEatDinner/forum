CREATE TABLE comments (
    comment_id   BIGINT AUTO_INCREMENT ,
    post_id      BIGINT NOT NULL,           -- 关联帖子ID
    author_id    BIGINT NOT NULL,           -- 关联用户ID
    content      LONGTEXT NOT NULL,         -- 评论内容
    create_time  DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_time DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`comment_id`),
    INDEX idx_post_id (`post_id`),
    INDEX idx_author_id (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;