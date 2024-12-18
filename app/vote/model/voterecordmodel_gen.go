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
	voteRecordFieldNames          = builder.RawFieldNames(&VoteRecord{})
	voteRecordRows                = strings.Join(voteRecordFieldNames, ",")
	voteRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(voteRecordFieldNames, "`vote_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	voteRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(voteRecordFieldNames, "`vote_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheVoteRecordVoteIdPrefix       = "cache:voteRecord:voteId:"
	cacheVoteRecordPostIdUserIdPrefix = "cache:voteRecord:postId:userId:"
)

type (
	voteRecordModel interface {
		FindOne(ctx context.Context, voteId int64) (*VoteRecord, error)
		FindOneByPostIdUserId(ctx context.Context, postId uint64, userId uint64) (*VoteRecord, error)
		Update(ctx context.Context, data *VoteRecord) error
		Delete(ctx context.Context, voteId int64) error
		CountUpvotes(ctx context.Context, postId uint64) (int64, error)
		CountDownvotes(ctx context.Context, postId uint64) (int64, error)
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error  // 开启事务
		Insert(ctx context.Context, session sqlx.Session, data *VoteRecord) (sql.Result, error) 
	}

	defaultVoteRecordModel struct {
		sqlc.CachedConn
		table string
	}

	VoteRecord struct {
		VoteId      int64     `db:"vote_id"`      // 投票记录ID
		PostId      uint64    `db:"post_id"`      // 帖子id
		UserId      uint64    `db:"user_id"`      // 用户ID
		VoteType    int64     `db:"vote_type"`    // 投票类型 0:赞成票 1:反对票
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdatedTime time.Time `db:"updated_time"` // 更新时间
	}
)

func newVoteRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultVoteRecordModel {
	return &defaultVoteRecordModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`vote_record`",
	}
}

func (m *defaultVoteRecordModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultVoteRecordModel) Delete(ctx context.Context, voteId int64) error {
	data, err := m.FindOne(ctx, voteId)
	if err != nil {
		return err
	}

	voteRecordPostIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheVoteRecordPostIdUserIdPrefix, data.PostId, data.UserId)
	voteRecordVoteIdKey := fmt.Sprintf("%s%v", cacheVoteRecordVoteIdPrefix, voteId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `vote_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, voteId)
	}, voteRecordPostIdUserIdKey, voteRecordVoteIdKey)
	return err
}

func (m *defaultVoteRecordModel) FindOne(ctx context.Context, voteId int64) (*VoteRecord, error) {
	voteRecordVoteIdKey := fmt.Sprintf("%s%v", cacheVoteRecordVoteIdPrefix, voteId)
	var resp VoteRecord
	err := m.QueryRowCtx(ctx, &resp, voteRecordVoteIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `vote_id` = ? limit 1", voteRecordRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, voteId)
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

func (m *defaultVoteRecordModel) FindOneByPostIdUserId(ctx context.Context, postId uint64, userId uint64) (*VoteRecord, error) {
	voteRecordPostIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheVoteRecordPostIdUserIdPrefix, postId, userId)
	var resp VoteRecord
	err := m.QueryRowIndexCtx(ctx, &resp, voteRecordPostIdUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `post_id` = ? and `user_id` = ? limit 1", voteRecordRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, postId, userId); err != nil {
			return nil, err
		}
		return resp.VoteId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultVoteRecordModel) Insert(ctx context.Context, session sqlx.Session, data *VoteRecord) (sql.Result, error) {
	voteRecordPostIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheVoteRecordPostIdUserIdPrefix, data.PostId, data.UserId)
	voteRecordVoteIdKey := fmt.Sprintf("%s%v", cacheVoteRecordVoteIdPrefix, data.VoteId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, voteRecordRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.PostId, data.UserId, data.VoteType, data.UpdatedTime)
		}
		return conn.ExecCtx(ctx, query, data.PostId, data.UserId, data.VoteType, data.UpdatedTime)
	}, voteRecordPostIdUserIdKey, voteRecordVoteIdKey)
}


func (m *defaultVoteRecordModel) Update(ctx context.Context, newData *VoteRecord) error {
	data, err := m.FindOne(ctx, newData.VoteId)
	if err != nil {
		return err
	}

	voteRecordPostIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheVoteRecordPostIdUserIdPrefix, data.PostId, data.UserId)
	voteRecordVoteIdKey := fmt.Sprintf("%s%v", cacheVoteRecordVoteIdPrefix, data.VoteId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `vote_id` = ?", m.table, voteRecordRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.PostId, newData.UserId, newData.VoteType, newData.UpdatedTime, newData.VoteId)
	}, voteRecordPostIdUserIdKey, voteRecordVoteIdKey)
	return err
}

func (m *defaultVoteRecordModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheVoteRecordVoteIdPrefix, primary)
}

func (m *defaultVoteRecordModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `vote_id` = ? limit 1", voteRecordRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVoteRecordModel) tableName() string {
	return m.table
}
// CountUpvotes 统计赞成票数量
func (m *defaultVoteRecordModel) CountUpvotes(ctx context.Context, postId uint64) (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE `post_id` = ? AND `vote_type` = 1", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, postId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountDownvotes 统计反对票数量
func (m *defaultVoteRecordModel) CountDownvotes(ctx context.Context, postId uint64) (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE `post_id` = ? AND `vote_type` = -1", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, postId)
	if err != nil {
		return 0, err
	}
	return count, nil
}
