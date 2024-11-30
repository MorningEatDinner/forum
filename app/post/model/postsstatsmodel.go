package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostsStatsModel = (*customPostsStatsModel)(nil)

type (
	// PostsStatsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsStatsModel.
	PostsStatsModel interface {
		postsStatsModel
	}

	customPostsStatsModel struct {
		*defaultPostsStatsModel
	}
)

// NewPostsStatsModel returns a model for the database table.
func NewPostsStatsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostsStatsModel {
	return &customPostsStatsModel{
		defaultPostsStatsModel: newPostsStatsModel(conn, c, opts...),
	}
}
