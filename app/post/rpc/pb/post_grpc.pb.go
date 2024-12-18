// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: app/post/rpc/pb/post.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PostService_CreatePost_FullMethodName             = "/pb.PostService/CreatePost"
	PostService_GetPostDetail_FullMethodName          = "/pb.PostService/GetPostDetail"
	PostService_GetPostList_FullMethodName            = "/pb.PostService/GetPostList"
	PostService_GetPostListByCommunity_FullMethodName = "/pb.PostService/GetPostListByCommunity"
	PostService_DeletePost_FullMethodName             = "/pb.PostService/DeletePost"
	PostService_UpdatePostScore_FullMethodName        = "/pb.PostService/UpdatePostScore"
	PostService_DeletePostScheduler_FullMethodName    = "/pb.PostService/DeletePostScheduler"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetPostDetail(ctx context.Context, in *GetPostDetailRequest, opts ...grpc.CallOption) (*GetPostDetailResponse, error)
	GetPostList(ctx context.Context, in *GetPostListRequest, opts ...grpc.CallOption) (*GetPostListResponse, error)
	GetPostListByCommunity(ctx context.Context, in *GetPostListByCommunityRequest, opts ...grpc.CallOption) (*GetPostListByCommunityResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	UpdatePostScore(ctx context.Context, in *UpdatePostScoreRequest, opts ...grpc.CallOption) (*UpdatePostScoreResponse, error)
	DeletePostScheduler(ctx context.Context, in *DeletePostSchedulerRequest, opts ...grpc.CallOption) (*DeletePostSchedulerResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostDetail(ctx context.Context, in *GetPostDetailRequest, opts ...grpc.CallOption) (*GetPostDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostDetailResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostList(ctx context.Context, in *GetPostListRequest, opts ...grpc.CallOption) (*GetPostListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostListResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostListByCommunity(ctx context.Context, in *GetPostListByCommunityRequest, opts ...grpc.CallOption) (*GetPostListByCommunityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostListByCommunityResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostListByCommunity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePostScore(ctx context.Context, in *UpdatePostScoreRequest, opts ...grpc.CallOption) (*UpdatePostScoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePostScoreResponse)
	err := c.cc.Invoke(ctx, PostService_UpdatePostScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePostScheduler(ctx context.Context, in *DeletePostSchedulerRequest, opts ...grpc.CallOption) (*DeletePostSchedulerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostSchedulerResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePostScheduler_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetPostDetail(context.Context, *GetPostDetailRequest) (*GetPostDetailResponse, error)
	GetPostList(context.Context, *GetPostListRequest) (*GetPostListResponse, error)
	GetPostListByCommunity(context.Context, *GetPostListByCommunityRequest) (*GetPostListByCommunityResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	UpdatePostScore(context.Context, *UpdatePostScoreRequest) (*UpdatePostScoreResponse, error)
	DeletePostScheduler(context.Context, *DeletePostSchedulerRequest) (*DeletePostSchedulerResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServiceServer struct{}

func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) GetPostDetail(context.Context, *GetPostDetailRequest) (*GetPostDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostDetail not implemented")
}
func (UnimplementedPostServiceServer) GetPostList(context.Context, *GetPostListRequest) (*GetPostListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostList not implemented")
}
func (UnimplementedPostServiceServer) GetPostListByCommunity(context.Context, *GetPostListByCommunityRequest) (*GetPostListByCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostListByCommunity not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) UpdatePostScore(context.Context, *UpdatePostScoreRequest) (*UpdatePostScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostScore not implemented")
}
func (UnimplementedPostServiceServer) DeletePostScheduler(context.Context, *DeletePostSchedulerRequest) (*DeletePostSchedulerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostScheduler not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}
func (UnimplementedPostServiceServer) testEmbeddedByValue()                     {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	// If the following call pancis, it indicates UnimplementedPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostDetail(ctx, req.(*GetPostDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostList(ctx, req.(*GetPostListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostListByCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostListByCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostListByCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostListByCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostListByCommunity(ctx, req.(*GetPostListByCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePostScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePostScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePostScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePostScore(ctx, req.(*UpdatePostScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePostScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostSchedulerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePostScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePostScheduler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePostScheduler(ctx, req.(*DeletePostSchedulerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPostDetail",
			Handler:    _PostService_GetPostDetail_Handler,
		},
		{
			MethodName: "GetPostList",
			Handler:    _PostService_GetPostList_Handler,
		},
		{
			MethodName: "GetPostListByCommunity",
			Handler:    _PostService_GetPostListByCommunity_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "UpdatePostScore",
			Handler:    _PostService_UpdatePostScore_Handler,
		},
		{
			MethodName: "DeletePostScheduler",
			Handler:    _PostService_DeletePostScheduler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/post/rpc/pb/post.proto",
}
