// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	commentsFieldNames          = builder.RawFieldNames(&Comments{})
	commentsRows                = strings.Join(commentsFieldNames, ",")
	commentsRowsExpectAutoSet   = strings.Join(stringx.Remove(commentsFieldNames, "`comment_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	commentsRowsWithPlaceHolder = strings.Join(stringx.Remove(commentsFieldNames, "`comment_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCommentsCommentIdPrefix = "cache:comments:commentId:"
)

type (
	commentsModel interface {
		Insert(ctx context.Context, data *Comments) (sql.Result, error)
		FindOne(ctx context.Context, commentId int64) (*Comments, error)
		Update(ctx context.Context, data *Comments) error
		Delete(ctx context.Context, commentId int64) error
		FindCommentListByPostId(ctx context.Context, postId int64, page, pageSize int64) ([]*Comments, error)
	}

	defaultCommentsModel struct {
		sqlc.CachedConn
		table string
	}

	Comments struct {
		CommentId   int64     `db:"comment_id"`
		PostId      int64     `db:"post_id"`
		AuthorId    int64     `db:"author_id"`
		Content     string    `db:"content"`
		CreateTime  time.Time `db:"create_time"`
		UpdatedTime time.Time `db:"updated_time"`
	}
)

func newCommentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultCommentsModel {
	return &defaultCommentsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`comments`",
	}
}

func (m *defaultCommentsModel) Delete(ctx context.Context, commentId int64) error {
	commentsCommentIdKey := fmt.Sprintf("%s%v", cacheCommentsCommentIdPrefix, commentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `comment_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, commentId)
	}, commentsCommentIdKey)
	return err
}

func (m *defaultCommentsModel) FindOne(ctx context.Context, commentId int64) (*Comments, error) {
	commentsCommentIdKey := fmt.Sprintf("%s%v", cacheCommentsCommentIdPrefix, commentId)
	var resp Comments
	err := m.QueryRowCtx(ctx, &resp, commentsCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", commentsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, commentId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentsModel) Insert(ctx context.Context, data *Comments) (sql.Result, error) {
	commentsCommentIdKey := fmt.Sprintf("%s%v", cacheCommentsCommentIdPrefix, data.CommentId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, commentsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.PostId, data.AuthorId, data.Content, data.UpdatedTime)
	}, commentsCommentIdKey)
	return ret, err
}

func (m *defaultCommentsModel) Update(ctx context.Context, data *Comments) error {
	commentsCommentIdKey := fmt.Sprintf("%s%v", cacheCommentsCommentIdPrefix, data.CommentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `comment_id` = ?", m.table, commentsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.PostId, data.AuthorId, data.Content, data.UpdatedTime, data.CommentId)
	}, commentsCommentIdKey)
	return err
}

func (m *defaultCommentsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCommentsCommentIdPrefix, primary)
}

func (m *defaultCommentsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", commentsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCommentsModel) tableName() string {
	return m.table
}



func (m *defaultCommentsModel) FindCommentListByPostId(ctx context.Context, postId int64, page, pageSize int64) ([]*Comments, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	var comments []*Comments
	query := fmt.Sprintf("select %s from %s where post_id = ? order by create_time desc limit ?, ?",
		commentsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &comments, query, postId, offset, pageSize)

	switch err {
	case nil:
		return comments, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
