// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package postservice

import (
	"context"

	"forum/tmp/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePostRequest              = pb.CreatePostRequest
	CreatePostResponse             = pb.CreatePostResponse
	DeletePostRequest              = pb.DeletePostRequest
	DeletePostResponse             = pb.DeletePostResponse
	GetPostDetailRequest           = pb.GetPostDetailRequest
	GetPostDetailResponse          = pb.GetPostDetailResponse
	GetPostListByCommunityRequest  = pb.GetPostListByCommunityRequest
	GetPostListByCommunityResponse = pb.GetPostListByCommunityResponse
	GetPostListRequest             = pb.GetPostListRequest
	GetPostListResponse            = pb.GetPostListResponse
	Post                           = pb.Post

	PostService interface {
		CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
		GetPostDetail(ctx context.Context, in *GetPostDetailRequest, opts ...grpc.CallOption) (*GetPostDetailResponse, error)
		GetPostList(ctx context.Context, in *GetPostListRequest, opts ...grpc.CallOption) (*GetPostListResponse, error)
		GetPostListByCommunity(ctx context.Context, in *GetPostListByCommunityRequest, opts ...grpc.CallOption) (*GetPostListByCommunityResponse, error)
		DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	}

	defaultPostService struct {
		cli zrpc.Client
	}
)

func NewPostService(cli zrpc.Client) PostService {
	return &defaultPostService{
		cli: cli,
	}
}

func (m *defaultPostService) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	client := pb.NewPostServiceClient(m.cli.Conn())
	return client.CreatePost(ctx, in, opts...)
}

func (m *defaultPostService) GetPostDetail(ctx context.Context, in *GetPostDetailRequest, opts ...grpc.CallOption) (*GetPostDetailResponse, error) {
	client := pb.NewPostServiceClient(m.cli.Conn())
	return client.GetPostDetail(ctx, in, opts...)
}

func (m *defaultPostService) GetPostList(ctx context.Context, in *GetPostListRequest, opts ...grpc.CallOption) (*GetPostListResponse, error) {
	client := pb.NewPostServiceClient(m.cli.Conn())
	return client.GetPostList(ctx, in, opts...)
}

func (m *defaultPostService) GetPostListByCommunity(ctx context.Context, in *GetPostListByCommunityRequest, opts ...grpc.CallOption) (*GetPostListByCommunityResponse, error) {
	client := pb.NewPostServiceClient(m.cli.Conn())
	return client.GetPostListByCommunity(ctx, in, opts...)
}

func (m *defaultPostService) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	client := pb.NewPostServiceClient(m.cli.Conn())
	return client.DeletePost(ctx, in, opts...)
}
