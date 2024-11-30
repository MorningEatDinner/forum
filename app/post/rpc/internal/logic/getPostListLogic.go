package logic

import (
	"cmp"
	"context"
	"forum/common/globalkey"
	"forum/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"slices"
	"strconv"

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
	// 1. 使用Redis计算得到post id list
	start, stop := (page-1)*size, page*size-1
	pids, err := l.svcCtx.RedisClient.ZrevrangeWithScores(globalkey.GetRedisKey(globalkey.PostScoreKey), start, stop)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("获取帖子列表失败"), "err:%v", err)
	}
	var postIds []int64
	postScores := make(map[int64]int64)
	for _, pid := range pids {
		id, _ := strconv.ParseInt(pid.Key, 10, 64)
		postIds = append(postIds, id) // 假设Redis返回的Member是int64类型
		postScores[id] = pid.Score
	}
	if len(postIds) == 0 {
		return []*model.Posts{}, nil
	}

	postList, err := l.postByIds(l.ctx, postIds)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("批量获取帖子失败, err: %v", err)
		return nil, err
	}
	// 还需要重新排序
	slices.SortFunc(postList, func(a, b *model.Posts) int {
		return cmp.Compare(postScores[b.PostId], postScores[a.PostId])
	})
	return postList, nil
}

func (l *GetPostListLogic) getPostListByCommunityId(id, page, size int64) ([]*model.Posts, error) {
	//list, err := l.svcCtx.PostModel.FindByCommunityId(l.ctx, id, page, size)
	//if err != nil {
	//	logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
	//	return nil, err
	//}
	//return list, nil
	return nil, nil
}

func (l *GetPostListLogic) getPostListByAuthorId(id, page, size int64) ([]*model.Posts, error) {
	//list, err := l.svcCtx.PostModel.FindByAuthorId(l.ctx, page, size, id)
	//if err != nil {
	//	logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
	//	return nil, err
	//}
	//return list, nil
	return nil, nil
}

func (l *GetPostListLogic) postByIds(ctx context.Context, postIds []int64) ([]*model.Posts, error) {
	posts, err := mr.MapReduce[int64, *model.Posts, []*model.Posts](
		// Map 阶段：发送有效的 postId 到源通道
		func(source chan<- int64) {
			for _, pid := range postIds {
				if pid == -1 {
					continue
				}
				source <- pid
			}
		},
		// Reduce 阶段：根据 postId 获取对应的帖子
		func(id int64, writer mr.Writer[*model.Posts], cancel func(error)) {
			post, err := l.svcCtx.PostModel.FindOne(l.ctx, id)
			if err != nil {
				// 如果查找失败，可以选择跳过或取消整个过程
				cancel(err)
				return
			}
			writer.Write(post)
		},
		// 聚合阶段：收集所有获取到的帖子
		func(pipe <-chan *model.Posts, writer mr.Writer[[]*model.Posts], cancel func(error)) {
			var posts []*model.Posts
			for post := range pipe {
				posts = append(posts, post)
			}
			writer.Write(posts)
		},
	)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post list failed, err: %v", err)
		return nil, err
	}

	return posts, nil
}
