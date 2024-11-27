package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VoteCountModel = (*customVoteCountModel)(nil)

type (
	// VoteCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVoteCountModel.
	VoteCountModel interface {
		voteCountModel
	}

	customVoteCountModel struct {
		*defaultVoteCountModel
	}
)

// NewVoteCountModel returns a model for the database table.
func NewVoteCountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VoteCountModel {
	return &customVoteCountModel{
		defaultVoteCountModel: newVoteCountModel(conn, c, opts...),
	}
}
