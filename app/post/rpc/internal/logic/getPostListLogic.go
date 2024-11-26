package logic

import (
	"context"

	"forum/app/post/model"
	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListLogic) GetPostList(in *pb.GetPostListRequest) (*pb.GetPostListResponse, error) {
	return l.getPostList0(in)
}

// 版本1 ： 没有引入分数， 而是直接按照创建时间进行排序
func (l *GetPostListLogic) getPostList0(in *pb.GetPostListRequest) (*pb.GetPostListResponse, error) {
	size := in.PageSize
	var list []*model.Posts
	var err error
	if in.PageSize > 10 {
		size = 10
	}
	if in.AuthorId == nil && in.CommunityId == nil {
		list, err = l.getPostListAll(in.Page, size)
	}
	if in.CommunityId != nil {
		list, err = l.getPostListByCommunityId(*in.CommunityId, in.Page, size)
	}
	if in.AuthorId != nil {
		list, err = l.getPostListByAuthorId(*in.AuthorId, in.Page, size)
	}
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, err
	}

	var respList []*pb.Post
	if len(list) > 0 {
		for _, v := range list {
			var modelPost pb.Post
			copier.Copy(&modelPost, v)
			respList = append(respList, &modelPost)
		}
	}

	return &pb.GetPostListResponse{
		Total: int64(len(respList)),
		Posts: respList,
	}, nil
}

func (l *GetPostListLogic) getPostListAll(page, size int64) ([]*model.Posts, error) {
	list, err := l.svcCtx.PostModel.FindAll(l.ctx, page, size)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, err
	}
	return list, nil
}

func (l *GetPostListLogic) getPostListByCommunityId(id, page, size int64) ([]*model.Posts, error) {
	list, err := l.svcCtx.PostModel.FindByCommunityId(l.ctx, id, page, size)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, err
	}
	return list, nil
}

func (l *GetPostListLogic) getPostListByAuthorId(id, page, size int64) ([]*model.Posts, error) {
	list, err := l.svcCtx.PostModel.FindByAuthorId(l.ctx, page, size, id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, err
	}
	return list, nil
}
