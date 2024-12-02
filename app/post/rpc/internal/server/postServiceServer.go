// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package server

import (
	"context"

	"forum/app/post/rpc/internal/logic"
	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"
)

type PostServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServer(svcCtx *svc.ServiceContext) *PostServiceServer {
	return &PostServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	l := logic.NewCreatePostLogic(ctx, s.svcCtx)
	return l.CreatePost(in)
}

func (s *PostServiceServer) GetPostDetail(ctx context.Context, in *pb.GetPostDetailRequest) (*pb.GetPostDetailResponse, error) {
	l := logic.NewGetPostDetailLogic(ctx, s.svcCtx)
	return l.GetPostDetail(in)
}

func (s *PostServiceServer) GetPostList(ctx context.Context, in *pb.GetPostListRequest) (*pb.GetPostListResponse, error) {
	l := logic.NewGetPostListLogic(ctx, s.svcCtx)
	return l.GetPostList(in)
}

func (s *PostServiceServer) GetPostListByCommunity(ctx context.Context, in *pb.GetPostListByCommunityRequest) (*pb.GetPostListByCommunityResponse, error) {
	l := logic.NewGetPostListByCommunityLogic(ctx, s.svcCtx)
	return l.GetPostListByCommunity(in)
}

func (s *PostServiceServer) DeletePost(ctx context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	l := logic.NewDeletePostLogic(ctx, s.svcCtx)
	return l.DeletePost(in)
}
func (s *PostServiceServer) UpdatePostScore(ctx context.Context, in *pb.UpdatePostScoreRequest) (*pb.UpdatePostScoreResponse, error) {
	l := logic.NewUpdatePostScoreLogic(ctx, s.svcCtx)
	return l.UpdatePostScore(in)
}

func (s *PostServiceServer) DeletePostScheduler(ctx context.Context, in *pb.DeletePostSchedulerRequest) (*pb.DeletePostSchedulerResponse, error) {
	l := logic.NewDeletePostSchedulerLogic(ctx, s.svcCtx)
	return l.DeletePostScheduler(in)
}
