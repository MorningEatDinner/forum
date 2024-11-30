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
	postsFieldNames          = builder.RawFieldNames(&Posts{})
	postsRows                = strings.Join(postsFieldNames, ",")
	postsRowsExpectAutoSet   = strings.Join(stringx.Remove(postsFieldNames, "`post_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	postsRowsWithPlaceHolder = strings.Join(stringx.Remove(postsFieldNames, "`post_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cachePostsPostIdPrefix = "cache:posts:postId:"
)

type (
	postsModel interface {
		Insert(ctx context.Context, data *Posts) (sql.Result, error)
		FindOne(ctx context.Context, postId int64) (*Posts, error)
		Update(ctx context.Context, data *Posts) error
		Delete(ctx context.Context, postId int64) error
	}

	defaultPostsModel struct {
		sqlc.CachedConn
		table string
	}

	Posts struct {
		PostId      int64     `db:"post_id"`      // 帖子ID
		AuthorId    int64     `db:"author_id"`    // 作者ID
		CommunityId int64     `db:"community_id"` // 社区ID
		Title       string    `db:"title"`        // 帖子标题
		Content     string    `db:"content"`      // 帖子内容
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdatedTime time.Time `db:"updated_time"` // 更新时间
	}
)

func newPostsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPostsModel {
	return &defaultPostsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`posts`",
	}
}

func (m *defaultPostsModel) Delete(ctx context.Context, postId int64) error {
	postsPostIdKey := fmt.Sprintf("%s%v", cachePostsPostIdPrefix, postId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `post_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, postId)
	}, postsPostIdKey)
	return err
}

func (m *defaultPostsModel) FindOne(ctx context.Context, postId int64) (*Posts, error) {
	postsPostIdKey := fmt.Sprintf("%s%v", cachePostsPostIdPrefix, postId)
	var resp Posts
	err := m.QueryRowCtx(ctx, &resp, postsPostIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `post_id` = ? limit 1", postsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, postId)
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

func (m *defaultPostsModel) Insert(ctx context.Context, data *Posts) (sql.Result, error) {
	postsPostIdKey := fmt.Sprintf("%s%v", cachePostsPostIdPrefix, data.PostId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, postsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AuthorId, data.CommunityId, data.Title, data.Content, data.UpdatedTime)
	}, postsPostIdKey)
	return ret, err
}

func (m *defaultPostsModel) Update(ctx context.Context, data *Posts) error {
	postsPostIdKey := fmt.Sprintf("%s%v", cachePostsPostIdPrefix, data.PostId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `post_id` = ?", m.table, postsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AuthorId, data.CommunityId, data.Title, data.Content, data.UpdatedTime, data.PostId)
	}, postsPostIdKey)
	return err
}

func (m *defaultPostsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePostsPostIdPrefix, primary)
}

func (m *defaultPostsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `post_id` = ? limit 1", postsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPostsModel) tableName() string {
	return m.table
}
