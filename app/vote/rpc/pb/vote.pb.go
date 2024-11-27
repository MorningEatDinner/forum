// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.19.4
// source: app/vote/rpc/pb/vote.proto

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

// 投票记录
type VoteRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteId      int64  `protobuf:"varint,1,opt,name=voteId,proto3" json:"voteId,omitempty"`          // 投票记录ID
	PostId      int64  `protobuf:"varint,2,opt,name=postId,proto3" json:"postId,omitempty"`          // 帖子ID
	UserId      int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`          // 用户ID
	VoteType    int32  `protobuf:"varint,4,opt,name=voteType,proto3" json:"voteType,omitempty"`      // 投票类型 -1: 反对票, 0: 取消投票, 1: 赞成票
	CreateTime  string `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`   // 创建时间
	UpdatedTime string `protobuf:"bytes,6,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"` // 更新时间
}

func (x *VoteRecord) Reset() {
	*x = VoteRecord{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRecord) ProtoMessage() {}

func (x *VoteRecord) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRecord.ProtoReflect.Descriptor instead.
func (*VoteRecord) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{0}
}

func (x *VoteRecord) GetVoteId() int64 {
	if x != nil {
		return x.VoteId
	}
	return 0
}

func (x *VoteRecord) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *VoteRecord) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VoteRecord) GetVoteType() int32 {
	if x != nil {
		return x.VoteType
	}
	return 0
}

func (x *VoteRecord) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *VoteRecord) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

// 投票计数
type VoteCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId       int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`             // 帖子ID
	AgreeCount   int32  `protobuf:"varint,2,opt,name=agreeCount,proto3" json:"agreeCount,omitempty"`     // 赞成票数
	OpposeCount  int32  `protobuf:"varint,3,opt,name=opposeCount,proto3" json:"opposeCount,omitempty"`   // 反对票数
	NeutralCount int32  `protobuf:"varint,4,opt,name=neutralCount,proto3" json:"neutralCount,omitempty"` // 取消票数 (如果需要的话)
	CreateTime   string `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`      // 创建时间
	UpdatedTime  string `protobuf:"bytes,6,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`    // 更新时间
}

func (x *VoteCount) Reset() {
	*x = VoteCount{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteCount) ProtoMessage() {}

func (x *VoteCount) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteCount.ProtoReflect.Descriptor instead.
func (*VoteCount) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{1}
}

func (x *VoteCount) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *VoteCount) GetAgreeCount() int32 {
	if x != nil {
		return x.AgreeCount
	}
	return 0
}

func (x *VoteCount) GetOpposeCount() int32 {
	if x != nil {
		return x.OpposeCount
	}
	return 0
}

func (x *VoteCount) GetNeutralCount() int32 {
	if x != nil {
		return x.NeutralCount
	}
	return 0
}

func (x *VoteCount) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *VoteCount) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

// 为帖子投票的请求
type VotePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64 `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`     // 帖子ID
	UserId   int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户ID
	VoteType int32 `protobuf:"varint,3,opt,name=voteType,proto3" json:"voteType,omitempty"` // 投票类型 -1: 反对票, 0: 取消投票, 1: 赞成票
}

func (x *VotePostRequest) Reset() {
	*x = VotePostRequest{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VotePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VotePostRequest) ProtoMessage() {}

func (x *VotePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VotePostRequest.ProtoReflect.Descriptor instead.
func (*VotePostRequest) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{2}
}

func (x *VotePostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *VotePostRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VotePostRequest) GetVoteType() int32 {
	if x != nil {
		return x.VoteType
	}
	return 0
}

// 为帖子投票的响应
type VotePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 投票是否成功
}

func (x *VotePostResponse) Reset() {
	*x = VotePostResponse{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VotePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VotePostResponse) ProtoMessage() {}

func (x *VotePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VotePostResponse.ProtoReflect.Descriptor instead.
func (*VotePostResponse) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{3}
}

func (x *VotePostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 更新用户投票的请求
type UpdateUserVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64 `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`     // 帖子ID
	UserId   int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户ID
	VoteType int32 `protobuf:"varint,3,opt,name=voteType,proto3" json:"voteType,omitempty"` // 新的投票类型 -1: 反对票, 0: 取消投票, 1: 赞成票
}

func (x *UpdateUserVoteRequest) Reset() {
	*x = UpdateUserVoteRequest{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserVoteRequest) ProtoMessage() {}

func (x *UpdateUserVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserVoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserVoteRequest) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserVoteRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UpdateUserVoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateUserVoteRequest) GetVoteType() int32 {
	if x != nil {
		return x.VoteType
	}
	return 0
}

// 更新用户投票的响应
type UpdateUserVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 更新投票是否成功
}

func (x *UpdateUserVoteResponse) Reset() {
	*x = UpdateUserVoteResponse{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserVoteResponse) ProtoMessage() {}

func (x *UpdateUserVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserVoteResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserVoteResponse) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserVoteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 撤销用户投票的请求
type RemoveUserVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteId int64 `protobuf:"varint,1,opt,name=voteId,proto3" json:"voteId,omitempty"` // 投票记录ID
}

func (x *RemoveUserVoteRequest) Reset() {
	*x = RemoveUserVoteRequest{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserVoteRequest) ProtoMessage() {}

func (x *RemoveUserVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserVoteRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserVoteRequest) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveUserVoteRequest) GetVoteId() int64 {
	if x != nil {
		return x.VoteId
	}
	return 0
}

// 撤销用户投票的响应
type RemoveUserVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 撤销投票是否成功
}

func (x *RemoveUserVoteResponse) Reset() {
	*x = RemoveUserVoteResponse{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserVoteResponse) ProtoMessage() {}

func (x *RemoveUserVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserVoteResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserVoteResponse) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveUserVoteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 获取帖子投票统计信息的请求
type GetVoteCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"` // 帖子ID
}

func (x *GetVoteCountRequest) Reset() {
	*x = GetVoteCountRequest{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoteCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteCountRequest) ProtoMessage() {}

func (x *GetVoteCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteCountRequest.ProtoReflect.Descriptor instead.
func (*GetVoteCountRequest) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{8}
}

func (x *GetVoteCountRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

// 获取帖子投票统计信息的响应
type GetVoteCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteCount *VoteCount `protobuf:"bytes,1,opt,name=voteCount,proto3" json:"voteCount,omitempty"` // 投票统计信息
}

func (x *GetVoteCountResponse) Reset() {
	*x = GetVoteCountResponse{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVoteCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteCountResponse) ProtoMessage() {}

func (x *GetVoteCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteCountResponse.ProtoReflect.Descriptor instead.
func (*GetVoteCountResponse) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{9}
}

func (x *GetVoteCountResponse) GetVoteCount() *VoteCount {
	if x != nil {
		return x.VoteCount
	}
	return nil
}

// 获取用户投票记录的请求
type GetUserVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"` // 帖子ID
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"` // 用户ID
}

func (x *GetUserVoteRequest) Reset() {
	*x = GetUserVoteRequest{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserVoteRequest) ProtoMessage() {}

func (x *GetUserVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserVoteRequest.ProtoReflect.Descriptor instead.
func (*GetUserVoteRequest) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserVoteRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *GetUserVoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 获取用户投票记录的响应
type GetUserVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteRecord *VoteRecord `protobuf:"bytes,1,opt,name=voteRecord,proto3" json:"voteRecord,omitempty"` // 用户的投票记录
}

func (x *GetUserVoteResponse) Reset() {
	*x = GetUserVoteResponse{}
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserVoteResponse) ProtoMessage() {}

func (x *GetUserVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_vote_rpc_pb_vote_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserVoteResponse.ProtoReflect.Descriptor instead.
func (*GetUserVoteResponse) Descriptor() ([]byte, []int) {
	return file_app_vote_rpc_pb_vote_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserVoteResponse) GetVoteRecord() *VoteRecord {
	if x != nil {
		return x.VoteRecord
	}
	return nil
}

var File_app_vote_rpc_pb_vote_proto protoreflect.FileDescriptor

var file_app_vote_rpc_pb_vote_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xb2, 0x01, 0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x61, 0x67, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f,
	0x70, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6f, 0x70, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x0f, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x63, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x15, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2d,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x32,
	0xe3, 0x02, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x08, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_vote_rpc_pb_vote_proto_rawDescOnce sync.Once
	file_app_vote_rpc_pb_vote_proto_rawDescData = file_app_vote_rpc_pb_vote_proto_rawDesc
)

func file_app_vote_rpc_pb_vote_proto_rawDescGZIP() []byte {
	file_app_vote_rpc_pb_vote_proto_rawDescOnce.Do(func() {
		file_app_vote_rpc_pb_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_vote_rpc_pb_vote_proto_rawDescData)
	})
	return file_app_vote_rpc_pb_vote_proto_rawDescData
}

var file_app_vote_rpc_pb_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_app_vote_rpc_pb_vote_proto_goTypes = []any{
	(*VoteRecord)(nil),             // 0: pb.VoteRecord
	(*VoteCount)(nil),              // 1: pb.VoteCount
	(*VotePostRequest)(nil),        // 2: pb.VotePostRequest
	(*VotePostResponse)(nil),       // 3: pb.VotePostResponse
	(*UpdateUserVoteRequest)(nil),  // 4: pb.UpdateUserVoteRequest
	(*UpdateUserVoteResponse)(nil), // 5: pb.UpdateUserVoteResponse
	(*RemoveUserVoteRequest)(nil),  // 6: pb.RemoveUserVoteRequest
	(*RemoveUserVoteResponse)(nil), // 7: pb.RemoveUserVoteResponse
	(*GetVoteCountRequest)(nil),    // 8: pb.GetVoteCountRequest
	(*GetVoteCountResponse)(nil),   // 9: pb.GetVoteCountResponse
	(*GetUserVoteRequest)(nil),     // 10: pb.GetUserVoteRequest
	(*GetUserVoteResponse)(nil),    // 11: pb.GetUserVoteResponse
}
var file_app_vote_rpc_pb_vote_proto_depIdxs = []int32{
	1,  // 0: pb.GetVoteCountResponse.voteCount:type_name -> pb.VoteCount
	0,  // 1: pb.GetUserVoteResponse.voteRecord:type_name -> pb.VoteRecord
	2,  // 2: pb.VoteService.VotePost:input_type -> pb.VotePostRequest
	4,  // 3: pb.VoteService.UpdateUserVote:input_type -> pb.UpdateUserVoteRequest
	6,  // 4: pb.VoteService.RemoveUserVote:input_type -> pb.RemoveUserVoteRequest
	8,  // 5: pb.VoteService.GetVoteCount:input_type -> pb.GetVoteCountRequest
	10, // 6: pb.VoteService.GetUserVote:input_type -> pb.GetUserVoteRequest
	3,  // 7: pb.VoteService.VotePost:output_type -> pb.VotePostResponse
	5,  // 8: pb.VoteService.UpdateUserVote:output_type -> pb.UpdateUserVoteResponse
	7,  // 9: pb.VoteService.RemoveUserVote:output_type -> pb.RemoveUserVoteResponse
	9,  // 10: pb.VoteService.GetVoteCount:output_type -> pb.GetVoteCountResponse
	11, // 11: pb.VoteService.GetUserVote:output_type -> pb.GetUserVoteResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_app_vote_rpc_pb_vote_proto_init() }
func file_app_vote_rpc_pb_vote_proto_init() {
	if File_app_vote_rpc_pb_vote_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_vote_rpc_pb_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_vote_rpc_pb_vote_proto_goTypes,
		DependencyIndexes: file_app_vote_rpc_pb_vote_proto_depIdxs,
		MessageInfos:      file_app_vote_rpc_pb_vote_proto_msgTypes,
	}.Build()
	File_app_vote_rpc_pb_vote_proto = out.File
	file_app_vote_rpc_pb_vote_proto_rawDesc = nil
	file_app_vote_rpc_pb_vote_proto_goTypes = nil
	file_app_vote_rpc_pb_vote_proto_depIdxs = nil
}