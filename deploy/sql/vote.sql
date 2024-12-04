create database vote;
use  vote;
-- auto-generated definition
create table vote_record
(
    vote_id      bigint auto_increment comment '投票记录ID'
        primary key,
    post_id      bigint unsigned default '0'                  not null comment '帖子id',
    user_id      bigint unsigned default '0'                  not null comment '用户ID',
    vote_type    tinyint         default 0                    not null comment '投票类型 0:赞成票 1:反对票',
    create_time  datetime(3)     default CURRENT_TIMESTAMP(3) not null comment '创建时间',
    updated_time datetime(3)     default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3) comment '更新时间',
    constraint uniq_post_user
        unique (post_id, user_id)
)
    comment '投票记录表' collate = utf8mb4_unicode_ci;

create index idx_user_id
    on vote_record (user_id);