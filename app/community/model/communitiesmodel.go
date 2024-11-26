package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommunitiesModel = (*customCommunitiesModel)(nil)

type (
	// CommunitiesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommunitiesModel.
	CommunitiesModel interface {
		communitiesModel
	}

	customCommunitiesModel struct {
		*defaultCommunitiesModel
	}
)

// NewCommunitiesModel returns a model for the database table.
func NewCommunitiesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommunitiesModel {
	return &customCommunitiesModel{
		defaultCommunitiesModel: newCommunitiesModel(conn, c, opts...),
	}
}
