// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.19.4
// source: app/comment/rpc/pb/comment.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 基础评论信息
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId   int64  `protobuf:"varint,1,opt,name=commentId,proto3" json:"commentId,omitempty"`    // 评论ID
	PostId      int64  `protobuf:"varint,2,opt,name=postId,proto3" json:"postId,omitempty"`          // 关联帖子ID
	AuthorId    int64  `protobuf:"varint,3,opt,name=authorId,proto3" json:"authorId,omitempty"`      // 关联用户ID
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`         // 评论内容
	CreateTime  string `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`   // 创建时间
	UpdatedTime string `protobuf:"bytes,6,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"` // 更新时间
}

func (x *Comment) Reset() {
	*x = Comment{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *Comment) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Comment) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Comment) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

// 创建评论请求
type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`     // 关联帖子ID
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`    // 评论内容
	AuthorId int64  `protobuf:"varint,3,opt,name=authorId,proto3" json:"authorId,omitempty"` // 作者ID（由API层通过JWT获取并传递）
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCommentRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CreateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateCommentRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

// 创建评论响应
type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"` // 创建的评论信息
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

// 删除评论请求
type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64 `protobuf:"varint,1,opt,name=commentId,proto3" json:"commentId,omitempty"` // 要删除的评论ID
	AuthorId  int64 `protobuf:"varint,2,opt,name=authorId,proto3" json:"authorId,omitempty"`   // 作者ID（由API层通过JWT获取并传递）
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCommentRequest) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *DeleteCommentRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

// 删除评论响应
type DeleteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{4}
}

// 根据帖子ID获取评论请求
type GetCommentsByPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64 `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`     // 要获取评论的帖子ID
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`         // 页码（可选，用于分页）
	PageSize int32 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"` // 每页数量（可选，用于分页）
}

func (x *GetCommentsByPostRequest) Reset() {
	*x = GetCommentsByPostRequest{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsByPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByPostRequest) ProtoMessage() {}

func (x *GetCommentsByPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByPostRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsByPostRequest) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentsByPostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *GetCommentsByPostRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCommentsByPostRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 根据帖子ID获取评论响应
type GetCommentsByPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`      // 评论总数
	Comments []*Comment `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"` // 获取到的评论列表
}

func (x *GetCommentsByPostResponse) Reset() {
	*x = GetCommentsByPostResponse{}
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsByPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByPostResponse) ProtoMessage() {}

func (x *GetCommentsByPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_comment_rpc_pb_comment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByPostResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsByPostResponse) Descriptor() ([]byte, []int) {
	return file_app_comment_rpc_pb_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetCommentsByPostResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetCommentsByPostResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

var File_app_comment_rpc_pb_comment_proto protoreflect.FileDescriptor

var file_app_comment_rpc_pb_comment_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb7, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x64, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x62, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x32, 0xf4, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_comment_rpc_pb_comment_proto_rawDescOnce sync.Once
	file_app_comment_rpc_pb_comment_proto_rawDescData = file_app_comment_rpc_pb_comment_proto_rawDesc
)

func file_app_comment_rpc_pb_comment_proto_rawDescGZIP() []byte {
	file_app_comment_rpc_pb_comment_proto_rawDescOnce.Do(func() {
		file_app_comment_rpc_pb_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_comment_rpc_pb_comment_proto_rawDescData)
	})
	return file_app_comment_rpc_pb_comment_proto_rawDescData
}

var file_app_comment_rpc_pb_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_comment_rpc_pb_comment_proto_goTypes = []any{
	(*Comment)(nil),                   // 0: pb.Comment
	(*CreateCommentRequest)(nil),      // 1: pb.CreateCommentRequest
	(*CreateCommentResponse)(nil),     // 2: pb.CreateCommentResponse
	(*DeleteCommentRequest)(nil),      // 3: pb.DeleteCommentRequest
	(*DeleteCommentResponse)(nil),     // 4: pb.DeleteCommentResponse
	(*GetCommentsByPostRequest)(nil),  // 5: pb.GetCommentsByPostRequest
	(*GetCommentsByPostResponse)(nil), // 6: pb.GetCommentsByPostResponse
}
var file_app_comment_rpc_pb_comment_proto_depIdxs = []int32{
	0, // 0: pb.CreateCommentResponse.comment:type_name -> pb.Comment
	0, // 1: pb.GetCommentsByPostResponse.comments:type_name -> pb.Comment
	1, // 2: pb.CommentService.CreateComment:input_type -> pb.CreateCommentRequest
	3, // 3: pb.CommentService.DeleteComment:input_type -> pb.DeleteCommentRequest
	5, // 4: pb.CommentService.GetCommentsByPost:input_type -> pb.GetCommentsByPostRequest
	2, // 5: pb.CommentService.CreateComment:output_type -> pb.CreateCommentResponse
	4, // 6: pb.CommentService.DeleteComment:output_type -> pb.DeleteCommentResponse
	6, // 7: pb.CommentService.GetCommentsByPost:output_type -> pb.GetCommentsByPostResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_comment_rpc_pb_comment_proto_init() }
func file_app_comment_rpc_pb_comment_proto_init() {
	if File_app_comment_rpc_pb_comment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_comment_rpc_pb_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_comment_rpc_pb_comment_proto_goTypes,
		DependencyIndexes: file_app_comment_rpc_pb_comment_proto_depIdxs,
		MessageInfos:      file_app_comment_rpc_pb_comment_proto_msgTypes,
	}.Build()
	File_app_comment_rpc_pb_comment_proto = out.File
	file_app_comment_rpc_pb_comment_proto_rawDesc = nil
	file_app_comment_rpc_pb_comment_proto_goTypes = nil
	file_app_comment_rpc_pb_comment_proto_depIdxs = nil
}
