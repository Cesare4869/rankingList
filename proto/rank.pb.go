// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: rank.proto

package proto

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

// 新增/更新 玩家分数
type UpdatePlayerRankInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid int64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
	Score  int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UpdatePlayerRankInfoReq) Reset() {
	*x = UpdatePlayerRankInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerRankInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerRankInfoReq) ProtoMessage() {}

func (x *UpdatePlayerRankInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerRankInfoReq.ProtoReflect.Descriptor instead.
func (*UpdatePlayerRankInfoReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePlayerRankInfoReq) GetRoleid() int64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

func (x *UpdatePlayerRankInfoReq) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type UpdatePlayerRankInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode int32 `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
}

func (x *UpdatePlayerRankInfoRes) Reset() {
	*x = UpdatePlayerRankInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerRankInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerRankInfoRes) ProtoMessage() {}

func (x *UpdatePlayerRankInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerRankInfoRes.ProtoReflect.Descriptor instead.
func (*UpdatePlayerRankInfoRes) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePlayerRankInfoRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

// 玩家信息
type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid int64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
	Score  int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Rank   int32 `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerInfo) GetRoleid() int64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

func (x *PlayerInfo) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *PlayerInfo) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

// 查询玩家分数
type QueryPlayerScoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid int64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
}

func (x *QueryPlayerScoreReq) Reset() {
	*x = QueryPlayerScoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlayerScoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlayerScoreReq) ProtoMessage() {}

func (x *QueryPlayerScoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlayerScoreReq.ProtoReflect.Descriptor instead.
func (*QueryPlayerScoreReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{3}
}

func (x *QueryPlayerScoreReq) GetRoleid() int64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

type QueryPlayerScoreRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	RetCode int32       `protobuf:"varint,2,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
}

func (x *QueryPlayerScoreRes) Reset() {
	*x = QueryPlayerScoreRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlayerScoreRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlayerScoreRes) ProtoMessage() {}

func (x *QueryPlayerScoreRes) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlayerScoreRes.ProtoReflect.Descriptor instead.
func (*QueryPlayerScoreRes) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{4}
}

func (x *QueryPlayerScoreRes) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *QueryPlayerScoreRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

// 查询玩家排名
type QueryPlayerRankReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid int64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
}

func (x *QueryPlayerRankReq) Reset() {
	*x = QueryPlayerRankReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlayerRankReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlayerRankReq) ProtoMessage() {}

func (x *QueryPlayerRankReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlayerRankReq.ProtoReflect.Descriptor instead.
func (*QueryPlayerRankReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{5}
}

func (x *QueryPlayerRankReq) GetRoleid() int64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

type QueryPlayerRankRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	RetCode int32       `protobuf:"varint,2,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
}

func (x *QueryPlayerRankRes) Reset() {
	*x = QueryPlayerRankRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPlayerRankRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlayerRankRes) ProtoMessage() {}

func (x *QueryPlayerRankRes) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlayerRankRes.ProtoReflect.Descriptor instead.
func (*QueryPlayerRankRes) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{6}
}

func (x *QueryPlayerRankRes) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *QueryPlayerRankRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

// 查询前五玩家
type QueryTop5RankReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryTop5RankReq) Reset() {
	*x = QueryTop5RankReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTop5RankReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTop5RankReq) ProtoMessage() {}

func (x *QueryTop5RankReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTop5RankReq.ProtoReflect.Descriptor instead.
func (*QueryTop5RankReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{7}
}

type QueryTop5RankRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    []*PlayerInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	RetCode int32         `protobuf:"varint,2,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
}

func (x *QueryTop5RankRes) Reset() {
	*x = QueryTop5RankRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTop5RankRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTop5RankRes) ProtoMessage() {}

func (x *QueryTop5RankRes) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTop5RankRes.ProtoReflect.Descriptor instead.
func (*QueryTop5RankRes) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{8}
}

func (x *QueryTop5RankRes) GetInfo() []*PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *QueryTop5RankRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

// 从榜上删除玩家数据
type DeletePlayerRankReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid int64 `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
}

func (x *DeletePlayerRankReq) Reset() {
	*x = DeletePlayerRankReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlayerRankReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlayerRankReq) ProtoMessage() {}

func (x *DeletePlayerRankReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlayerRankReq.ProtoReflect.Descriptor instead.
func (*DeletePlayerRankReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePlayerRankReq) GetRoleid() int64 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

type DeletePlayerRankRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode int32 `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
}

func (x *DeletePlayerRankRes) Reset() {
	*x = DeletePlayerRankRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlayerRankRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlayerRankRes) ProtoMessage() {}

func (x *DeletePlayerRankRes) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlayerRankRes.ProtoReflect.Descriptor instead.
func (*DeletePlayerRankRes) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePlayerRankRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

// 清空排行榜数据
type ClearRankInofReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearRankInofReq) Reset() {
	*x = ClearRankInofReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearRankInofReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRankInofReq) ProtoMessage() {}

func (x *ClearRankInofReq) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRankInofReq.ProtoReflect.Descriptor instead.
func (*ClearRankInofReq) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{11}
}

type ClearRankInofRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode int32 `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
}

func (x *ClearRankInofRes) Reset() {
	*x = ClearRankInofRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearRankInofRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRankInofRes) ProtoMessage() {}

func (x *ClearRankInofRes) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRankInofRes.ProtoReflect.Descriptor instead.
func (*ClearRankInofRes) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{12}
}

func (x *ClearRankInofRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

var File_rank_proto protoreflect.FileDescriptor

var file_rank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x22, 0x47, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x4e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x22, 0x2d, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x22,
	0x55, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x6f, 0x70, 0x35, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x22, 0x52,
	0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x70, 0x35, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69,
	0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49,
	0x6e, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x22, 0x2c, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52,
	0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rank_proto_rawDescOnce sync.Once
	file_rank_proto_rawDescData = file_rank_proto_rawDesc
)

func file_rank_proto_rawDescGZIP() []byte {
	file_rank_proto_rawDescOnce.Do(func() {
		file_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_rank_proto_rawDescData)
	})
	return file_rank_proto_rawDescData
}

var file_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_rank_proto_goTypes = []interface{}{
	(*UpdatePlayerRankInfoReq)(nil), // 0: rank.UpdatePlayerRankInfoReq
	(*UpdatePlayerRankInfoRes)(nil), // 1: rank.UpdatePlayerRankInfoRes
	(*PlayerInfo)(nil),              // 2: rank.PlayerInfo
	(*QueryPlayerScoreReq)(nil),     // 3: rank.QueryPlayerScoreReq
	(*QueryPlayerScoreRes)(nil),     // 4: rank.QueryPlayerScoreRes
	(*QueryPlayerRankReq)(nil),      // 5: rank.QueryPlayerRankReq
	(*QueryPlayerRankRes)(nil),      // 6: rank.QueryPlayerRankRes
	(*QueryTop5RankReq)(nil),        // 7: rank.QueryTop5RankReq
	(*QueryTop5RankRes)(nil),        // 8: rank.QueryTop5RankRes
	(*DeletePlayerRankReq)(nil),     // 9: rank.DeletePlayerRankReq
	(*DeletePlayerRankRes)(nil),     // 10: rank.DeletePlayerRankRes
	(*ClearRankInofReq)(nil),        // 11: rank.ClearRankInofReq
	(*ClearRankInofRes)(nil),        // 12: rank.ClearRankInofRes
}
var file_rank_proto_depIdxs = []int32{
	2, // 0: rank.QueryPlayerScoreRes.info:type_name -> rank.PlayerInfo
	2, // 1: rank.QueryPlayerRankRes.info:type_name -> rank.PlayerInfo
	2, // 2: rank.QueryTop5RankRes.info:type_name -> rank.PlayerInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rank_proto_init() }
func file_rank_proto_init() {
	if File_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerRankInfoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerRankInfoRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlayerScoreReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlayerScoreRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlayerRankReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPlayerRankRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTop5RankReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTop5RankRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlayerRankReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlayerRankRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearRankInofReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rank_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearRankInofRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rank_proto_goTypes,
		DependencyIndexes: file_rank_proto_depIdxs,
		MessageInfos:      file_rank_proto_msgTypes,
	}.Build()
	File_rank_proto = out.File
	file_rank_proto_rawDesc = nil
	file_rank_proto_goTypes = nil
	file_rank_proto_depIdxs = nil
}
